package handler

import "github.com/FelipeSoft/traffik-one/internal/core/usecase"

type Container struct {
	TestHandler         *TestHandler
	BackendHandler      *BackendHandler
	RoutingRulesHandler *RoutingRulesHandler
}

func NewContainer(usecaseContainer *usecase.Container) *Container {
	return &Container{
		TestHandler:         NewTestHandler(usecaseContainer.TestUseCase),
		BackendHandler:      NewBackendHandler(usecaseContainer.BackendUseCase),
		RoutingRulesHandler: NewRoutingRulesHandler(usecaseContainer.RoutingRulesUseCase),
	}
}
