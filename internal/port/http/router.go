package http

import (
	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

func RegisterRoutes(app *app.App) *port.Router {
	router := port.NewRouter()

	router.Handle("POST", "/test", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.TestHandler.Test()))
	router.Handle("POST", "/backends/add", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.AddBackend()))
	router.Handle("GET", "/backends", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.GetAllBackends()))
	router.Handle("PUT", "/backends/{backendId}/update", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.UpdateBackend()))
	router.Handle("PATCH", "/backends/{backendId}/activate", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.ActivateBackend()))
	router.Handle("PATCH", "/backends/{backendId}/inactivate", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.InactivateBackend()))
	router.Handle("DELETE", "/backends/{backendId}/delete", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.DeleteBackend()))
	router.Handle("GET", "/backends/{backendId}/find", app.Middlewares.AuthenticationMiddleware.HandleBearerToken(app.Handlers.BackendHandler.GetBackendByID()))

	return router
}
