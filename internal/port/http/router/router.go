package router

import (
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/app"
)

func NewHttpRouter(app *app.App) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/test", app.Handlers.TestHandler.Test())
	mux.Handle("/backends/add", app.Handlers.BackendHandler.AddBackend())
	return mux
}
