package usecase

import (
	"context"
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type TestUseCase struct {
	repo port.TestRepository
}

func NewTestUseCase(repo port.TestRepository) *TestUseCase {
	return &TestUseCase{
		repo: repo,
	}
}

func (uc *TestUseCase) Test(ctx context.Context, message string) (string, error) {
	if message != "Ping" {
		return "", fmt.Errorf(`please provide the "Ping" as message param`)
	}
	return "Pong", nil
}
