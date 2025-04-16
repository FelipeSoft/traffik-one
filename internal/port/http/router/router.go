package router

import (
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/app"
)

func NewHttpRouter(app *app.App) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/test", app.Handlers.TestHandler.Test())
	return mux
}
