package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/core/usecase"
)

type RoutingRulesHandler struct {
	uc *usecase.RoutingRulesUseCase
}

func NewRoutingRulesHandler(uc *usecase.RoutingRulesUseCase) *RoutingRulesHandler {
	return &RoutingRulesHandler{
		uc: uc,
	}
}

func (h *RoutingRulesHandler) AddRoutingRules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		var dto dto.AddRoutingRulesInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Invalid request body: %v", err)))
			return
		}
		
		if dto.Source == "" || dto.Target == "" || dto.Protocol == "" || dto.PoolID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing required fields"))
			return
		}

		err := h.uc.Add(ctx, dto)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *RoutingRulesHandler) UpdateRoutingRules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		routingRulesId := params["routingRulesId"]
		if routingRulesId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'backendId' missing"))
			return
		}

		var dto dto.UpdateRoutingRulesInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(fmt.Appendf(nil, "Invalid request body: %v", err))
			return
		}

		dto.ID = routingRulesId

		err := h.uc.Update(ctx, dto)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *RoutingRulesHandler) GetRoutingRulesByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		routingRulesId := params["routingRulesId"]
		if routingRulesId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'routingRulesId' missing"))
			return
		}

		backends, err := h.uc.GetRoutingRulesById(ctx, dto.GetRoutingRulesByIDInput{
			ID: routingRulesId,
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

func (h *RoutingRulesHandler) GetRoutingRulesByPoolID() http.HandlerFunc {
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

		backends, err := h.uc.GetRoutingRulesByPoolID(ctx, dto.GetRoutingRulesByPoolIDInput{
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

func (h *RoutingRulesHandler) GetAllRoutingRules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		backends, err := h.uc.GetAllRoutingRules(ctx)
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

func (h *RoutingRulesHandler) DeleteRoutingRules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, ok := ctx.Value(port.ParamsKey).(map[string]string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid context params"))
			return
		}

		routingRulesId := params["routingRulesId"]
		if routingRulesId == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The URL param 'routingRulesId' missing"))
			return
		}

		err := h.uc.Delete(ctx, dto.DeleteRoutingRulesInput{
			ID: routingRulesId,
		})
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
