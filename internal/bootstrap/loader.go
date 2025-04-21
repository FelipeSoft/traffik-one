package bootstrap

import (
	"context"
	"log"
	"os"

	"github.com/FelipeSoft/traffik-one/internal/adapter/repository"
	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
)

func LoadInitialConfig() *entity.ConfigEvent {
	ctx := context.Background()

	poolId := os.Getenv("POOL_ID")
	db := bolt.DB()

	backendRepository := repository.NewBoltBackendRepository(db)
	routingRulesRepository := repository.NewBoltRoutingRulesRepository(db)
	algorithmsRepository := repository.NewBoltAlgorithmsRepository(db)

	backends, err := backendRepository.FindBackendsByPoolID(ctx, poolId, true)
	if err != nil || len(backends) == 0 {
		log.Println("[Bootstrap] No backends found, using empty list")
		backends = []entity.Backend{}
	}

	routingRules, err := routingRulesRepository.FindRoutingRulesByPoolID(ctx, poolId)
	if err != nil || len(routingRules) == 0 {
		log.Println("[Bootstrap] No routing rules found, using empty list")
		routingRules = []entity.RoutingRules{}
	}

	algorithm, err := algorithmsRepository.Get(ctx)
	if err != nil && algorithm == "" {
		log.Println("[Bootstrap] No algorithm found, using default RoundRobin")
		defaultAlgorithm := "wrr"
		if saveErr := algorithmsRepository.Set(ctx, defaultAlgorithm); saveErr != nil {
			log.Fatalf("failed to save default algorithm: %v", saveErr)
		}
		algorithm = defaultAlgorithm
	}

	return &entity.ConfigEvent{
		Backend:      backends,
		RoutingRules: routingRules,
		Algorithm:    &algorithm,
	}
}

func Load() *entity.ConfigEvent {
	return nil
}
