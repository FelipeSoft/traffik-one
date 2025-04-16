package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
)

type TestHandler struct {
	uc  *usecase.TestUseCase
}

func NewTestHandler(uc *usecase.TestUseCase) *TestHandler {
	return &TestHandler{
		uc: uc,
	}
}

func (h *TestHandler) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var test entity.Test
		if err := json.Unmarshal(body, &test); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid JSON: %v", err))
			return
		}

		message, err := h.uc.Test(ctx, test.Message)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Could not execute the usecase.TestUseCase: %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	}
}
