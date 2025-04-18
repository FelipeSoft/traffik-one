package handler

import (
	"encoding/json"
	"fmt"
	"io"
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

		err = h.uc.AddBackend(ctx, dto)
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
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		var dto dto.UpdateBackendInput
		err = json.Unmarshal(body, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

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

		dto.ID = backendId

		err = h.uc.UpdateBackend(ctx, dto)
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
