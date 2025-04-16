package repository

import "context"

type MemoryBackendRepository struct {
	backends []any
}

func NewMemoryBackendRepository() *MemoryBackendRepository {
	return &MemoryBackendRepository{}
}

func (r *MemoryBackendRepository) Save(ctx context.Context) error {
	r.backends = append(r.backends, "")
	return nil
}