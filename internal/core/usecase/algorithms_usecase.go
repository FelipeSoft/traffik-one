package usecase

import (
	"context"
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/port/dispatcher"
)

type AlgorithmsUseCase struct {
	repo       port.AlgorithmsRepository
	dispatcher *dispatcher.AlgorithmsDispatcher
}

func NewAlgorithmsUseCase(repo port.AlgorithmsRepository, dispatcher *dispatcher.AlgorithmsDispatcher) *AlgorithmsUseCase {
	return &AlgorithmsUseCase{
		repo:       repo,
		dispatcher: dispatcher,
	}
}

func (uc *AlgorithmsUseCase) Set(ctx context.Context, providedAlgorithm string) error {
	allowedAlgorithms := []string{"crr", "wrr", "lc0"}
	includesOnAlgorithmsList := false

	for _, algorithm := range allowedAlgorithms {
		if providedAlgorithm == algorithm {
			includesOnAlgorithmsList = true
		}
	}

	notIncludesOnAlgorithmsList := !includesOnAlgorithmsList
	if notIncludesOnAlgorithmsList {
		return fmt.Errorf("could not use the provided algorithm")
	}

	currentAlgorithm, err := uc.repo.Get(ctx)
	if err != nil {
		return err
	}

	if currentAlgorithm == providedAlgorithm {
		return fmt.Errorf("the provided algorithm already is setted")
	}

	err = uc.repo.Set(ctx, providedAlgorithm)
	if err != nil {
		return err
	}

	dispatchAlgorithm, err := uc.repo.Get(ctx)
	if err != nil {
		return err
	}

	uc.dispatcher.Dispatch(dispatchAlgorithm)
	return nil
}

func (uc *AlgorithmsUseCase) Get(ctx context.Context) (string, error) {
	algorithm, err := uc.repo.Get(ctx)
	if err != nil {
		return "", err
	}
	return algorithm, nil
}
