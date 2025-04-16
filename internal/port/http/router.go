package http

import (
	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

func RegisterRoutes(app *app.App) *port.Router {
	router := port.NewRouter()

	router.Handle("POST", "/test", app.Middlewares.AuthenticationMiddleware.HandleAuth(app.Handlers.TestHandler.Test()))
	router.Handle("POST", "/backends/add", app.Middlewares.AuthenticationMiddleware.HandleAuth(app.Handlers.BackendHandler.AddBackend()))
	router.Handle("GET", "/backends", app.Middlewares.AuthenticationMiddleware.HandleAuth(app.Handlers.BackendHandler.GetAllBackends()))
	router.Handle("POST", "/backends/{backendId}/remove/pool", app.Handlers.BackendHandler.RemoveBackendFromPool())
	router.Handle("POST", "/backends/update/{backendId}", app.Handlers.BackendHandler.UpdateBackend())
	router.Handle("POST", "/backends/{backendId}/activate", app.Handlers.BackendHandler.ActivateBackend())
	router.Handle("POST", "/backends/{backendId}/inactivate", app.Handlers.BackendHandler.InactivateBackend())
	router.Handle("DELETE", "/backends/delete/{backendId}", app.Handlers.BackendHandler.DeleteBackend())
	router.Handle("GET", "/backends/{backendId}/find", app.Handlers.BackendHandler.GetBackendByID())

	return router
}
