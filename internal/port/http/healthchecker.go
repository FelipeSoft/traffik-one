package http

import "github.com/FelipeSoft/traffik-one/internal/core/entity"

type HttpHealthChecker struct {
	configEvent *entity.ConfigEvent
}

func NewHealthChecker(configEvent *entity.ConfigEvent) *HttpHealthChecker {
	return &HttpHealthChecker{
		configEvent: configEvent,
	}
}

func (h *HttpHealthChecker) StartHttpHealthChecker() {
	
}
