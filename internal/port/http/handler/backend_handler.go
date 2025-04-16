package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
)

type BackendHandler struct {
	uc *usecase.BackendUseCase
}

func NewBackendHandler(uc *usecase.BackendUseCase) *BackendHandler {
	return &BackendHandler{
		uc: uc,
	}
}

func (h *BackendHandler) AddBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}

		_, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Backend Added succesfully!"))
	}
}
