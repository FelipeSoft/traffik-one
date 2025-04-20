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
var backendToPoolBoltBucket []byte = []byte("backend_by_pool")

func NewBoltBackendRepository(db *bolt.DB) *BoltBackendRepository {
	return &BoltBackendRepository{
		db: db,
	}
}

func (r *BoltBackendRepository) Save(ctx context.Context, backend *entity.Backend) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		mainBucket := tx.Bucket(backendBoltBucket)
		if mainBucket == nil {
			return fmt.Errorf("bucket 'backends' not found")
		}

		existingRaw := mainBucket.Get([]byte(backend.ID))

		var previous entity.Backend
		var previousPoolID string

		if existingRaw != nil {
			if err := json.Unmarshal(existingRaw, &previous); err == nil {
				previousPoolID = previous.PoolID
			}
		}

		if backend.PoolID == "" && previousPoolID != "" {
			backend.PoolID = previousPoolID
		}

		marshalBackend, err := json.Marshal(backend)
		if err != nil {
			return err
		}

		err = mainBucket.Put([]byte(backend.ID), marshalBackend)
		if err != nil {
			return err
		}

		indexPoolBucket := tx.Bucket(backendToPoolBoltBucket)
		if indexPoolBucket == nil {
			return fmt.Errorf("bucket 'backend_by_pool' not found")
		}

		if previousPoolID != "" && previousPoolID != backend.PoolID {
			oldSub := indexPoolBucket.Bucket([]byte(previousPoolID))
			if oldSub != nil {
				_ = oldSub.Delete([]byte(backend.ID))
			}
		}

		if backend.PoolID == "" {
			return fmt.Errorf("cannot index backend without poolId")
		}

		newSub, err := indexPoolBucket.CreateBucketIfNotExists([]byte(backend.PoolID))
		if err != nil {
			return err
		}
		err = newSub.Put([]byte(backend.ID), marshalBackend)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *BoltBackendRepository) Delete(ctx context.Context, backendId string, poolId string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(backendBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(backendBoltBucket))
		}
		
		err := bucket.Delete([]byte(backendId))
		if err != nil {
			return err
		}

		indexBucket := tx.Bucket([]byte(backendToPoolBoltBucket))
		subBucket := indexBucket.Bucket([]byte(poolId))
		err = subBucket.Delete([]byte(backendId))
		if err != nil {
			return err
		}
		return nil
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
	var result entity.Backend
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(backendBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(backendBoltBucket))
		}
		value := bucket.Get([]byte(backendId))
		if len(value) == 0 {
			return fmt.Errorf("could not find the backend with id '%s'", backendId)
		}
		err := json.Unmarshal(value, &result)
		return err
	})
	return &result, err
}

func (r *BoltBackendRepository) FindBackendsByPoolID(ctx context.Context, poolId string) ([]entity.Backend, error) {
	var results []entity.Backend

	err := r.db.View(func(tx *bolt.Tx) error {
		backendsByPoolsBucket := tx.Bucket([]byte(backendToPoolBoltBucket))
		mainBucket := tx.Bucket(backendBoltBucket)

		if backendsByPoolsBucket == nil {
			return fmt.Errorf("required buckets 'backend_by_pool' not found")
		}

		if mainBucket == nil {
			return fmt.Errorf("required buckets 'backends' not found")
		}

		indexValue := backendsByPoolsBucket.Bucket([]byte(poolId))
		if indexValue == nil {
			return fmt.Errorf("no backends found for poolId %s", poolId)
		}
		err := indexValue.ForEach(func (k []byte, v []byte) error {
			var backend entity.Backend
			if err := json.Unmarshal(v, &backend); err != nil {
				return err
			}
			results = append(results, backend)
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})

	return results, err
}
