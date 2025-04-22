package http

import (
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/core/port/websocket"
)

func RegisterRoutes(app *app.App, ws *websocket.WebsocketServer) *port.Router {
	router := port.NewRouter()

	// withBearerToken := app.Middlewares.AuthenticationMiddleware.HandleBearerToken

	router.Handle("GET", "/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWebsocket(ws, w, r)
	})
	router.Handle("POST", "/test", app.Handlers.TestHandler.Test())

	router.Handle("POST", "/backends/add", app.Handlers.BackendHandler.AddBackend())
	router.Handle("GET", "/backends", app.Handlers.BackendHandler.GetAllBackends())
	router.Handle("PUT", "/backends/{backendId}/update", app.Handlers.BackendHandler.UpdateBackend())
	router.Handle("PATCH", "/backends/{backendId}/activate", app.Handlers.BackendHandler.ActivateBackend())
	router.Handle("PATCH", "/backends/{backendId}/inactivate", app.Handlers.BackendHandler.InactivateBackend())
	router.Handle("DELETE", "/backends/{backendId}/delete", app.Handlers.BackendHandler.DeleteBackend())
	router.Handle("GET", "/backends/{backendId}/find", app.Handlers.BackendHandler.GetBackendByID())
	router.Handle("GET", "/backends/pool/{poolId}/find", app.Handlers.BackendHandler.GetBackendByPoolID())

	router.Handle("POST", "/routing/rules/add", app.Handlers.RoutingRulesHandler.AddRoutingRules())
	router.Handle("PUT", "/routing/rules/{routingRulesId}/update", app.Handlers.RoutingRulesHandler.UpdateRoutingRules())
	router.Handle("GET", "/routing/rules/{routingRulesId}/find", app.Handlers.RoutingRulesHandler.GetRoutingRulesByID())
	router.Handle("GET", "/routing/rules", app.Handlers.RoutingRulesHandler.GetAllRoutingRules())
	router.Handle("DELETE", "/routing/rules/{routingRulesId}/delete", app.Handlers.RoutingRulesHandler.DeleteRoutingRules())
	router.Handle("GET", "/routing/rules/pool/{poolId}/find", app.Handlers.RoutingRulesHandler.GetRoutingRulesByPoolID())

	router.Handle("GET", "/algorithm/get", app.Handlers.AlgorithmsHandler.GetAlgorithm())
	router.Handle("PUT", "/algorithm/set", app.Handlers.AlgorithmsHandler.SetAlgorithm())

	return router
}
