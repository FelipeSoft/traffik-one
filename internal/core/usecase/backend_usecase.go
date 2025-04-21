package usecase

import (
	"context"
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/port/dispatcher"
)

type BackendUseCase struct {
	repo       port.BackendRepository
	dispatcher *dispatcher.BackendDispatcher
}

func NewBackendUseCase(repo port.BackendRepository, dispatcher *dispatcher.BackendDispatcher) *BackendUseCase {
	return &BackendUseCase{
		repo:       repo,
		dispatcher: dispatcher,
	}
}

func (uc *BackendUseCase) AddBackend(ctx context.Context, input dto.AddBackendInput) error {
	if input.PoolID != "1" {
		return fmt.Errorf("only the default poolId 1 should be used")
	}

	backend, err := entity.NewBackend(
		input.IPv4,
		input.Hostname,
		input.Port,
		input.Protocol,
		input.Weight,
		input.PoolID,
	)
	if err != nil {
		return err
	}

	err = uc.repo.Save(ctx, backend)
	if err != nil {
		return err
	}

	backends, err := uc.repo.FindBackendsByPoolID(ctx, input.PoolID, true)
	if err != nil {
		return err
	}

	uc.dispatcher.Dispatch(backends)
	return nil
}

func (uc *BackendUseCase) UpdateBackend(ctx context.Context, input dto.UpdateBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	updatedBackend, err := backend.Update(
		input.IPv4,
		input.Hostname,
		input.Port,
		input.Protocol,
		input.Weight,
		input.PoolID,
	)
	if err != nil {
		return err
	}

	err = uc.repo.Save(ctx, updatedBackend)
	if err != nil {
		return err
	}

	backends, err := uc.repo.FindBackendsByPoolID(ctx, backend.PoolID, true)
	if err != nil {
		return err
	}

	uc.dispatcher.Dispatch(backends)
	return nil
}

func (uc *BackendUseCase) ActivateBackend(ctx context.Context, input dto.ActivateBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	if err = backend.Activate(); err != nil {
		return err
	}
	if err = uc.repo.Save(ctx, backend); err != nil {
		return err
	}

	backends, err := uc.repo.FindBackendsByPoolID(ctx, backend.PoolID, true)
	if err != nil {
		return err
	}
	uc.dispatcher.Dispatch(backends)
	return nil
}

func (uc *BackendUseCase) InactivateBackend(ctx context.Context, input dto.InactivateBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	if err = backend.Inactivate(); err != nil {
		return err
	}
	if err = uc.repo.Save(ctx, backend); err != nil {
		return err
	}
	backends, err := uc.repo.FindBackendsByPoolID(ctx, backend.PoolID, true)
	if err != nil {
		return err
	}
	uc.dispatcher.Dispatch(backends)
	return nil
}

func (uc *BackendUseCase) DeleteBackend(ctx context.Context, input dto.DeleteBackendInput) error {
	backend, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	err = uc.repo.Delete(ctx, input.ID, backend.PoolID)
	if err != nil {
		return err
	}
	backends, err := uc.repo.FindBackendsByPoolID(ctx, backend.PoolID, true)
	if err != nil {
		return err
	}
	uc.dispatcher.Dispatch(backends)
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

func (uc *BackendUseCase) GetBackendsByPoolID(ctx context.Context, input dto.GetBackendsByPoolIDInput) ([]entity.Backend, error) {
	backends, err := uc.repo.FindBackendsByPoolID(ctx, input.PoolID, false)
	if err != nil {
		return backends, err
	}
	return backends, nil
}
