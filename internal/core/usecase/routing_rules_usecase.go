package usecase

import (
	"context"
	"time"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type RoutingRulesUseCase struct {
	repo port.RoutingRulesRepository
}

func NewRoutingRulesUseCase(repo port.RoutingRulesRepository) *RoutingRulesUseCase {
	return &RoutingRulesUseCase{
		repo: repo,
	}
}

func (uc *RoutingRulesUseCase) Add(ctx context.Context, input dto.AddRoutingRulesInput) error {
	routingRules := entity.NewRoutingRules(
		input.Source,
		input.Target,
		input.Protocol,
	)

	return uc.repo.Save(ctx, routingRules)
}

func (uc *RoutingRulesUseCase) Update(ctx context.Context, input dto.UpdateRoutingRulesInput) error {
	routingRules, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	routingRules.Target = input.Target
	routingRules.Source = input.Source
	routingRules.Protocol = input.Protocol
	routingRules.UpdatedAt = time.Now()

	err = uc.repo.Save(ctx, routingRules)
	if err != nil {
		return err
	}
	return nil
}

func (uc *RoutingRulesUseCase) Delete(ctx context.Context, input dto.DeleteRoutingRulesInput) error {
	return uc.repo.Delete(ctx, input.ID)
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
