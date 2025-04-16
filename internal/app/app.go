package app

import (
	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
	"github.com/FelipeSoft/traffik-one/internal/port/http/handler"
)

type App struct {
	Handlers *handler.Container
	UseCases *usecase.Container
}

func NewApp() *App {
	usecases := usecase.NewContainer() 
	handlers := handler.NewContainer(usecases)

	return &App{
		Handlers: handlers,
		UseCases: usecases,
	}
}