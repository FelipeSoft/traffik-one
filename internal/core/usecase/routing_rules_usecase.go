package usecase

import (
	"context"
	"time"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/port/dispatcher"
)

type RoutingRulesUseCase struct {
	repo       port.RoutingRulesRepository
	dispatcher *dispatcher.RoutingRulesDispatcher
}

func NewRoutingRulesUseCase(repo port.RoutingRulesRepository, dispatcher *dispatcher.RoutingRulesDispatcher) *RoutingRulesUseCase {
	return &RoutingRulesUseCase{
		repo:       repo,
		dispatcher: dispatcher,
	}
}

func (uc *RoutingRulesUseCase) Add(ctx context.Context, input dto.AddRoutingRulesInput) error {
	routingRules := entity.NewRoutingRules(
		input.Source,
		input.Target,
		input.Protocol,
		input.PoolID,
	)

	err := uc.repo.Save(ctx, routingRules)
	if err != nil {
		return err
	}

	currentRoutingRules, err := uc.repo.FindRoutingRulesByPoolID(ctx, input.PoolID)
	if err != nil {
		return err
	}

	uc.dispatcher.Dispatch(currentRoutingRules)
	return nil
}

func (uc *RoutingRulesUseCase) Update(ctx context.Context, input dto.UpdateRoutingRulesInput) error {
	routingRules, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if input.Target != "" {
		routingRules.Target = input.Target
	}

	if input.Source != "" {
		routingRules.Source = input.Source
	}

	if input.Protocol != "" {
		routingRules.Protocol = input.Protocol
	}

	if input.PoolID != "" {
		routingRules.PoolID = input.PoolID
	}

	routingRules.UpdatedAt = time.Now()

	err = uc.repo.Save(ctx, routingRules)
	if err != nil {
		return err
	}

	currentRoutingRules, err := uc.repo.FindRoutingRulesByPoolID(ctx, routingRules.PoolID)
	if err != nil {
		return err
	}

	uc.dispatcher.Dispatch(currentRoutingRules)
	return nil
}

func (uc *RoutingRulesUseCase) Delete(ctx context.Context, input dto.DeleteRoutingRulesInput) error {
	existsRoutingRules, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	err = uc.repo.Delete(ctx, input.ID, existsRoutingRules.PoolID)
	if err != nil {
		return err
	}
	dispatchRoutingRules, err := uc.repo.FindRoutingRulesByPoolID(ctx, existsRoutingRules.PoolID)
	if err != nil {
		return err
	}
	uc.dispatcher.Dispatch(dispatchRoutingRules)
	return nil
}

func (uc *RoutingRulesUseCase) GetAllRoutingRules(ctx context.Context) ([]entity.RoutingRules, error) {
	routingRules, err := uc.repo.GetAll(ctx)
	if err != nil {
		return routingRules, err
	}
	return routingRules, nil
}

func (uc *RoutingRulesUseCase) GetRoutingRulesById(ctx context.Context, input dto.GetRoutingRulesByIDInput) (*entity.RoutingRules, error) {
	routingRules, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return routingRules, err
	}
	return routingRules, nil
}

func (uc *RoutingRulesUseCase) GetRoutingRulesByPoolID(ctx context.Context, input dto.GetRoutingRulesByPoolIDInput) ([]entity.RoutingRules, error) {
	routingRules, err := uc.repo.FindRoutingRulesByPoolID(ctx, input.PoolID)
	if err != nil {
		return routingRules, err
	}
	return routingRules, nil
}
