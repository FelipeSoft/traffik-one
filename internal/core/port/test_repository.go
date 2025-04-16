package port

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type TestRepository interface {
	ResponseThePingCommand(ctx context.Context, test entity.Test) string
}
