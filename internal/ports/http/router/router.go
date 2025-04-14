package router

import (
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/ports/http/handler"
)

func NewHttpRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler.TestHandler)
	return mux
}
