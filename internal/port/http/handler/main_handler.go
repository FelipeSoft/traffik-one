package handler

import (
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type MainHandler struct {
	algorithm port.Algorithm
}

func NewMainHandler(algorithm port.Algorithm) *MainHandler {
	return &MainHandler{
		algorithm: algorithm,
	}
}

func (h *MainHandler) HandleReverseProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.algorithm.ReverseProxy(w, r)
	}
}