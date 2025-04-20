package app

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
	"github.com/FelipeSoft/traffik-one/internal/port/algorithm"
	"github.com/FelipeSoft/traffik-one/internal/port/http/handler"
	"github.com/FelipeSoft/traffik-one/internal/port/http/middleware"
)

type App struct {
	Handlers    *handler.Container
	UseCases    *usecase.Container
	Middlewares *middleware.Container
	Algorithms  *algorithm.Container
}

func NewApp(ctx context.Context, configEvent *entity.ConfigEvent) *App {
	usecases := usecase.NewContainer(ctx, configEvent)
	handlers := handler.NewContainer(usecases)
	middlewares := middleware.NewContainer()
	algorithms := algorithm.NewContainer(configEvent)

	return &App{
		Handlers:    handlers,
		UseCases:    usecases,
		Middlewares: middlewares,
		Algorithms:  algorithms,
	}
}
