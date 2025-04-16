package bolt

import (
	"log"
	"os"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	dbInstance *bolt.DB
	once       sync.Once
)

func Init(path string, database string) {
	once.Do(func() {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, 0666)
			if err != nil {
				log.Fatalf("Could not create the database folder: %v", err)
			}
		}

		conn, err := bolt.Open(path + database, 0666, nil)
		if err != nil {
			log.Fatalf("Error to open BoltDB connection: %v", err)
		}
		dbInstance = conn
	})
}

func DB() *bolt.DB {
	if dbInstance == nil {
		log.Fatal("BoltDB was not initialized. Call bolt.Init() first.")
	}
	return dbInstance
}

func Close() {
	if dbInstance != nil {
		_ = dbInstance.Close()
	}
}
