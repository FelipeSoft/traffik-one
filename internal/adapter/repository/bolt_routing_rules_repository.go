package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/boltdb/bolt"
)

type BoltRoutingRulesRepository struct {
	db *bolt.DB
}

var routingRulesBoltBucket []byte = []byte("routing_rules")

func NewBoltRoutingRulesRepository(db *bolt.DB) *BoltRoutingRulesRepository {
	return &BoltRoutingRulesRepository{
		db: db,
	}
}

func (r *BoltRoutingRulesRepository) Save(ctx context.Context, routingRules *entity.RoutingRules) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(routingRulesBoltBucket)
		marshalBackend, err := json.Marshal(routingRules)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(routingRules.ID), marshalBackend)
		return err
	})
}

func (r *BoltRoutingRulesRepository) Delete(ctx context.Context, routingRulesId string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(routingRulesBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(routingRulesBoltBucket))
		}
		err := bucket.Delete([]byte(routingRulesId))
		return err
	})
}

func (r *BoltRoutingRulesRepository) GetAll(ctx context.Context) ([]entity.RoutingRules, error) {
	var result []entity.RoutingRules
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(routingRulesBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(routingRulesBoltBucket))
		}
		return bucket.ForEach(func(k []byte, v []byte) error {
			var routingRules entity.RoutingRules
			err := json.Unmarshal(v, &routingRules)
			if err != nil {
				return err
			}
			result = append(result, routingRules)
			return nil
		})
	})
	return result, err
}

func (r *BoltRoutingRulesRepository) GetByID(ctx context.Context, routingRulesId string) (*entity.RoutingRules, error) {
	var result entity.RoutingRules
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(routingRulesBoltBucket)
		if bucket == nil {
			return fmt.Errorf("bolt bucket named as '%s' not found", string(routingRulesBoltBucket))
		}
		value := bucket.Get([]byte(routingRulesId))
		if len(value) == 0 {
			return fmt.Errorf("could not find the backend with id '%s'", routingRulesId)
		}
		err := json.Unmarshal(value, &result)
		return err
	})
	return &result, err
}
