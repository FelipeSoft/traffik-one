package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/boltdb/bolt"
)

type BoltBackendRepository struct {
	db *bolt.DB
}

var backendBoltBucket []byte = []byte("backends")

func NewBoltBackendRepository(db *bolt.DB) *BoltBackendRepository {
	return &BoltBackendRepository{
		db: db,
	}
}

func (r *BoltBackendRepository) Save(ctx context.Context, backend *entity.Backend) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(backendBoltBucket)
		marshalBackend, err := json.Marshal(backend)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(backend.ID), marshalBackend)
		return err
	})
}

func (r *BoltBackendRepository) Delete(ctx context.Context, backendId string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(backendBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(backendBoltBucket))
		}
		err := bucket.Delete([]byte(backendId))
		return err
	})
}

func (r *BoltBackendRepository) GetAll(ctx context.Context) ([]entity.Backend, error) {
	var result []entity.Backend
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(backendBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(backendBoltBucket))
		}
		return bucket.ForEach(func(k []byte, v []byte) error {
			var backend entity.Backend
			err := json.Unmarshal(v, &backend)
			if err != nil {
				return err
			}
			result = append(result, backend)
			return nil
		})
	})
	return result, err
}

func (r *BoltBackendRepository) GetByID(ctx context.Context, backendId string) (*entity.Backend, error) {
	var result *entity.Backend
	r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(backendBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(backendBoltBucket))
		}
		value := bucket.Get([]byte(backendId))
		err := json.Unmarshal(value, &result)
		return err
	})
	return result, nil
}
