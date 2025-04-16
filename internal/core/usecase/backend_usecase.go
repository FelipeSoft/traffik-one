package usecase

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/entity"
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

func (uc *BackendUseCase) AddBackend(ctx context.Context, input dto.AddBackendInput) error {
	backend := &entity.Backend{
		IPv4:     input.IPv4,
		Hostname: input.Hostname,
		Port:     input.Port,
		Protocol: input.Protocol,
		Weight:   input.Weight,
		PoolID:   &input.PoolID,
	}
	backend.Activate()
	uc.repo.Save(ctx, backend)
	return nil
}

func (uc *BackendUseCase) RemoveBackendFromPool(ctx context.Context, input dto.RemoveBackendFromPoolInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	backend.RemoveFromPool()
	err = uc.repo.Save(ctx, backend)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BackendUseCase) UpdateBackend(ctx context.Context, input dto.UpdateBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	backend.Hostname = input.Hostname
	backend.IPv4 = input.IPv4
	backend.PoolID = &input.PoolID
	backend.Port = input.Port
	backend.Protocol = input.Protocol
	backend.Weight = input.Weight

	err = uc.repo.Save(ctx, backend)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BackendUseCase) ActivateBackend(ctx context.Context, input dto.ActivateBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	backend.Activate()
	err = uc.repo.Save(ctx, backend)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BackendUseCase) InactivateBackend(ctx context.Context, input dto.InactivateBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	backend.Inactivate()
	err = uc.repo.Save(ctx, backend)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BackendUseCase) DeleteBackend(ctx context.Context, input dto.InactivateBackendInput) error {
	err := uc.repo.Delete(ctx, input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BackendUseCase) GetAllBackends(ctx context.Context) ([]entity.Backend, error) {
	backends, err := uc.repo.GetAll(ctx)
	if err != nil {
		return backends, err
	}
	return backends, nil
}

func (uc *BackendUseCase) GetBackendByID(ctx context.Context, input dto.GetBackendByIDInput) (*entity.Backend, error) {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return backend, err
	}
	return backend, nil
}
