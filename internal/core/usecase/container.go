package usecase

import "github.com/FelipeSoft/traffik-one/internal/adapter/repository"

type Container struct {
	TestUseCase *TestUseCase
	// new usecases could be placed here...
}

func NewContainer() *Container {
	// injecting the repository adapter
	testRepository := repository.NewMemoryTestRepository()

	return &Container{
		TestUseCase: NewTestUseCase(testRepository), // create the usecase injecting the repository
		// new usecases could be placed here...
	}
}