package http

import (
	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

func RegisterRoutes(app *app.App) *port.Router {
	router := port.NewRouter()

	withBearerToken := app.Middlewares.AuthenticationMiddleware.HandleBearerToken

	router.Handle("POST", "/test",
		withBearerToken(app.Handlers.TestHandler.Test()))
	router.Handle("POST", "/backends/add",
		withBearerToken(app.Handlers.BackendHandler.AddBackend()))
	router.Handle("GET", "/backends",
		withBearerToken(app.Handlers.BackendHandler.GetAllBackends()))
	router.Handle("PUT", "/backends/{backendId}/update",
		withBearerToken(app.Handlers.BackendHandler.UpdateBackend()))
	router.Handle("PATCH", "/backends/{backendId}/activate",
		withBearerToken(app.Handlers.BackendHandler.ActivateBackend()))
	router.Handle("PATCH", "/backends/{backendId}/inactivate",
		withBearerToken(app.Handlers.BackendHandler.InactivateBackend()))
	router.Handle("DELETE", "/backends/{backendId}/delete",
		withBearerToken(app.Handlers.BackendHandler.DeleteBackend()))
	router.Handle("GET", "/backends/{backendId}/find",
		withBearerToken(app.Handlers.BackendHandler.GetBackendByID()))

	return router
}
