package usecase

import (
	"context"
	"log"

	"github.com/FelipeSoft/traffik-one/internal/adapter/repository"
	"github.com/FelipeSoft/traffik-one/internal/bootstrap"
	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
	"github.com/FelipeSoft/traffik-one/internal/port/dispatcher"
)

type Container struct {
	ConfigEvent         *entity.ConfigEvent
	TestUseCase         *TestUseCase
	BackendUseCase      *BackendUseCase
	RoutingRulesUseCase *RoutingRulesUseCase
	AlgorithmsUseCase   *AlgorithmsUseCase
}

func NewContainer(ctx context.Context) *Container {
	boltDb := bolt.DB()
	configEvent := bootstrap.LoadInitialConfig()

	backendDispatcher := dispatcher.NewBackendDispatcher(configEvent, make(chan []entity.Backend, 1))
	routingRulesDispatcher := dispatcher.NewRoutingRulesDispatcher(configEvent, make(chan []entity.RoutingRules, 1))
	algorithmsDispatcher := dispatcher.NewAlgorithmsDispatcher(configEvent, make(chan string, 1))

	dispatcher := dispatcher.NewDispatcher(configEvent, []port.Dispatcher{
		backendDispatcher,
		routingRulesDispatcher,
		algorithmsDispatcher,
	})
	dispatcher.Start(ctx)

	testRepository := repository.NewMemoryTestRepository()
	backendRepository := repository.NewBoltBackendRepository(boltDb)
	routingRulesRepository := repository.NewBoltRoutingRulesRepository(boltDb)
	algorithmsRepository := repository.NewBoltAlgorithmsRepository(boltDb)

	testUseCase := NewTestUseCase(testRepository)
	backendUseCase := NewBackendUseCase(backendRepository, backendDispatcher)
	routingRulesUseCase := NewRoutingRulesUseCase(routingRulesRepository, routingRulesDispatcher)
	algorithmsRulesUseCase := NewAlgorithmsUseCase(algorithmsRepository, algorithmsDispatcher)

	log.Println("[UseCase Container] Dependencies loaded successfully")

	return &Container{
		ConfigEvent:         configEvent,
		TestUseCase:         testUseCase,
		BackendUseCase:      backendUseCase,
		RoutingRulesUseCase: routingRulesUseCase,
		AlgorithmsUseCase:   algorithmsRulesUseCase,
	}
}
