package repository

import (
	"context"

	"github.com/boltdb/bolt"
)

type BoltAlgorithmsRepository struct {
	db *bolt.DB
}

func NewBoltAlgorithmsRepository(db *bolt.DB) *BoltAlgorithmsRepository {
	return &BoltAlgorithmsRepository{
		db: db,
	}
}

func (r *BoltAlgorithmsRepository) Get(ctx context.Context) (string, error) {
	var algorithm []byte

	err := r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("current_algorithm"))
		algorithm = bucket.Get([]byte("root"))
		return nil
	})

	return string(algorithm), err
}

func (r *BoltAlgorithmsRepository) Set(ctx context.Context, algorithm string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("current_algorithm"))
		err := bucket.Put([]byte("root"), []byte(algorithm))
		if err != nil {
			return err
		}
		return nil
	})
}
