package bolt

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	dbInstance *bolt.DB
	once       sync.Once
	initError  error
)

func Init(path string, database string) error {
	once.Do(func() {
		if err := os.MkdirAll(path, 0755); err != nil {
			initError = err
			return
		}

		dbPath := filepath.Join(path, database)
		conn, err := bolt.Open(dbPath, 0600, nil)
		if err != nil {
			initError = err
			return
		}

		bucketNames := []string{"backends", "pools", "routing_rules", "current_algorithm", "backend_by_pool", "routing_rules_by_pool"}
		if err := startBuckets(conn, bucketNames); err != nil {
			conn.Close()
			initError = err
			return
		}

		dbInstance = conn
		err = dbInstance.Update(func(tx *bolt.Tx) error {
			currentAlgorithm := tx.Bucket([]byte("current_algorithm"))
			algorithm := currentAlgorithm.Get([]byte("root"))
			if algorithm == nil {
				err := currentAlgorithm.Put([]byte("root"), []byte("wrr"))
				if err != nil {
					return err
				}
			}

			return err
		})
		initError = err
	})

	return initError
}

func startBuckets(db *bolt.DB, bucketNames []string) error {
	return db.Update(func(tx *bolt.Tx) error {
		for _, name := range bucketNames {
			if _, err := tx.CreateBucketIfNotExists([]byte(name)); err != nil {
				return err
			}
			log.Printf("[BoltDB] Bucket '%s' initilized with success", name)
		}
		return nil
	})
}

func DB() *bolt.DB {
	if dbInstance == nil {
		log.Fatal("[BoltDB] A instance was not initialized. Call bolt.Init() first.")
	}
	return dbInstance
}

func Close() {
	if dbInstance != nil {
		if err := dbInstance.Close(); err != nil {
			log.Printf("[BoltDB] Error on close connection with BoltDB: %v", err)
		}
		dbInstance = nil
	}
}
