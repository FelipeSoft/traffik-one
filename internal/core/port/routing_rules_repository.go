package port

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type RoutingRulesRepository interface {
	Save(ctx context.Context, routingRules *entity.RoutingRules) error
	Delete(ctx context.Context, routingRulesId string) error
	GetAll(ctx context.Context) ([]entity.RoutingRules, error)
	GetByID(ctx context.Context, routingRulesId string) (*entity.RoutingRules, error)
}
