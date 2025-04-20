package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
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
		defer r.Body.Close()
		ctx := r.Context()

		var dto dto.AddBackendInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		if dto.Hostname == "" || dto.IPv4 == "" || dto.PoolID == "" || dto.Port == 0 || dto.Weight == 0 || dto.Protocol == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing required fields"))
			return
		}

		err := h.uc.AddBackend(ctx, dto)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) UpdateBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		backendId := params["backendId"]
		if backendId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'backendId' missing"))
			return
		}

		var dto dto.UpdateBackendInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		dto.ID = backendId

		err := h.uc.UpdateBackend(ctx, dto)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *BackendHandler) ActivateBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		backendId := params["backendId"]
		if backendId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'backendId' missing"))
			return
		}

		err := h.uc.ActivateBackend(ctx, dto.ActivateBackendInput{
			ID: backendId,
		})
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) InactivateBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		backendId := params["backendId"]
		if backendId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'backendId' missing"))
			return
		}

		err := h.uc.InactivateBackend(ctx, dto.InactivateBackendInput{
			ID: backendId,
		})
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *BackendHandler) DeleteBackend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		backendId := params["backendId"]
		if backendId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'backendId' missing"))
			return
		}

		err := h.uc.DeleteBackend(ctx, dto.DeleteBackendInput{
			ID: backendId,
		})
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *BackendHandler) GetAllBackends() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		backends, err := h.uc.GetAllBackends(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := json.Marshal(backends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (h *BackendHandler) GetBackendByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		backendId := params["backendId"]
		if backendId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'backendId' missing"))
			return
		}

		backends, err := h.uc.GetBackendByID(ctx, dto.GetBackendByIDInput{
			ID: backendId,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := json.Marshal(backends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (h *BackendHandler) GetBackendByPoolID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		poolId := params["poolId"]
		if poolId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'poolId' missing"))
			return
		}

		backends, err := h.uc.GetBackendsByPoolID(ctx, dto.GetBackendsByPoolIDInput{
			PoolID: poolId,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := json.Marshal(backends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
