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

	router.Handle("POST", "/routing/rules/add",
		withBearerToken(app.Handlers.RoutingRulesHandler.AddRoutingRules()))
	router.Handle("PUT", "/routing/rules/{routingRulesId}/update",
		withBearerToken(app.Handlers.RoutingRulesHandler.UpdateRoutingRules()))
	router.Handle("GET", "/routing/rules/{routingRulesId}/find",
		withBearerToken(app.Handlers.RoutingRulesHandler.GetRoutingRulesByID()))
	router.Handle("GET", "/routing/rules",
		withBearerToken(app.Handlers.RoutingRulesHandler.GetAllRoutingRules()))
	router.Handle("DELETE", "/routing/rules/{routingRulesId}/delete",
		withBearerToken(app.Handlers.RoutingRulesHandler.DeleteRoutingRules()))

	return router
}
