package usecase

import (
	"github.com/FelipeSoft/traffik-one/internal/adapter/repository"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
)

type Container struct {
	TestUseCase    *TestUseCase
	BackendUseCase *BackendUseCase
}

func NewContainer() *Container {
	boltDb := bolt.DB()

	testRepository := repository.NewMemoryTestRepository()
	backendRepository := repository.NewBoltBackendRepository(boltDb)

	return &Container{
		TestUseCase:    NewTestUseCase(testRepository),
		BackendUseCase: NewBackendUseCase(backendRepository),
	}
}
