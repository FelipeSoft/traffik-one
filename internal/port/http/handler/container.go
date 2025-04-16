package handler

import "github.com/FelipeSoft/traffik-one/internal/core/usecase"

type Container struct {
	TestHandler *TestHandler
	// new http handlers could be placed here...
}

func NewContainer(usecaseContainer *usecase.Container) *Container {
	return &Container{
		TestHandler: NewTestHandler(usecaseContainer.TestUseCase), // create the http handler injecting the usecase
		// new http handlers could be placed here...
	}
}