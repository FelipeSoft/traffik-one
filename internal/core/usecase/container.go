package usecase

import (
	"github.com/FelipeSoft/traffik-one/internal/adapter/repository"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
)

type Container struct {
	TestUseCase         *TestUseCase
	BackendUseCase      *BackendUseCase
	RoutingRulesUseCase *RoutingRulesUseCase
}

func NewContainer() *Container {
	boltDb := bolt.DB()

	testRepository := repository.NewMemoryTestRepository()
	backendRepository := repository.NewBoltBackendRepository(boltDb)
	routingRulesRepository := repository.NewBoltRoutingRulesRepository(boltDb)

	return &Container{
		TestUseCase:         NewTestUseCase(testRepository),
		BackendUseCase:      NewBackendUseCase(backendRepository),
		RoutingRulesUseCase: NewRoutingRulesUseCase(routingRulesRepository),
	}
}
