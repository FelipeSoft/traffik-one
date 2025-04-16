package usecase

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type BackendUseCase struct {
	repo port.BackendRepository
}

func NewBackendUseCase(repo port.BackendRepository) *BackendUseCase {
	return &BackendUseCase{
		repo: repo,
	}
}

func (uc *BackendUseCase) AddBackend(ctx context.Context) {

}
