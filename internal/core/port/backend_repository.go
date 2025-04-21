package port

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type BackendRepository interface {
	Save(ctx context.Context, backend *entity.Backend) error
	GetAll(ctx context.Context) ([]entity.Backend, error)
	GetByID(ctx context.Context, backendId string) (*entity.Backend, error)
	Delete(ctx context.Context, backendId string, poolId string) error
	FindBackendsByPoolID(ctx context.Context, poolId string, availables bool) ([]entity.Backend, error)
}
