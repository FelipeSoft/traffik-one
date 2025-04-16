package app

import (
	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
	"github.com/FelipeSoft/traffik-one/internal/port/http/handler"
	"github.com/FelipeSoft/traffik-one/internal/port/http/middleware"
)

type App struct {
	Handlers    *handler.Container
	UseCases    *usecase.Container
	Middlewares *middleware.Container
}

func NewApp() *App {
	usecases := usecase.NewContainer()
	handlers := handler.NewContainer(usecases)
	middlewares := middleware.NewContainer()

	return &App{
		Handlers:    handlers,
		UseCases:    usecases,
		Middlewares: middlewares,
	}
}
