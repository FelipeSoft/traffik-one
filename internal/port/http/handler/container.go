package handler

import (
	"log"

	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
)

type Container struct {
	TestHandler         *TestHandler
	BackendHandler      *BackendHandler
	RoutingRulesHandler *RoutingRulesHandler
	AlgorithmsHandler   *AlgorithmsHandler
}

func NewContainer(usecaseContainer *usecase.Container) *Container {
	log.Println("[Handler Container] Dependencies loaded successfully")

	return &Container{
		TestHandler:         NewTestHandler(usecaseContainer.TestUseCase),
		BackendHandler:      NewBackendHandler(usecaseContainer.BackendUseCase),
		RoutingRulesHandler: NewRoutingRulesHandler(usecaseContainer.RoutingRulesUseCase),
		AlgorithmsHandler:   NewAlgorithmsHandler(usecaseContainer.AlgorithmsUseCase),
	}
}
