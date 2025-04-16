package usecase

import "github.com/FelipeSoft/traffik-one/internal/adapter/repository"

type Container struct {
	TestUseCase    *TestUseCase
	BackendUseCase *BackendUseCase
}

func NewContainer() *Container {
	testRepository := repository.NewMemoryTestRepository()
	backendRepository := repository.NewMemoryBackendRepository()

	return &Container{
		TestUseCase:    NewTestUseCase(testRepository),
		BackendUseCase: NewBackendUseCase(backendRepository),
	}
}
