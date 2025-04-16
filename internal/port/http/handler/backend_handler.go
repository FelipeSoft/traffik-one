package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
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

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) RemoveBackendFromPool() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) UpdateBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) ActivateBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) InactivateBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) DeleteBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) GetAllBackends() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		backends, err := h.uc.GetAllBackends(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		resp, err := json.Marshal(backends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (h *BackendHandler) GetBackendByID() http.HandlerFunc {
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

		var dto dto.AddBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		h.uc.AddBackend(ctx, dto)
		w.WriteHeader(http.StatusCreated)
	}
}
