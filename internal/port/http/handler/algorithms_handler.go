package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
)

type AlgorithmsHandler struct {
	uc *usecase.AlgorithmsUseCase
}

func NewAlgorithmsHandler(uc *usecase.AlgorithmsUseCase) *AlgorithmsHandler {
	return &AlgorithmsHandler{
		uc: uc,
	}
}

func (h *AlgorithmsHandler) SetAlgorithm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}
		defer r.Body.Close()

		var dto dto.SetAlgorithmInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		err = h.uc.Set(ctx, dto.Algorithm)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AlgorithmsHandler) GetAlgorithm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		algorithm, err := h.uc.Get(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := json.Marshal(algorithm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
