package repository

import (
	"context"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type MemoryTestRepository struct {
	tests []entity.Test
}

func NewMemoryTestRepository() *MemoryTestRepository {
	return &MemoryTestRepository{}
}

func (r *MemoryTestRepository) ResponseThePingCommand(ctx context.Context, test entity.Test) string {
	r.tests = append(r.tests, test)
	return "Pong"
}
