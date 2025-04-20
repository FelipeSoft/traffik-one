package algorithm

import (
	"log"
	"net/http"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type WeightedRoundRobinAlgorithm struct {
	repo     port.BackendRepository
	backends []entity.Backend
	index    uint32
	mu       sync.RWMutex
}

func NewWeightedRoundRobinAlgorithm(repo port.BackendRepository) *WeightedRoundRobinAlgorithm {
	return &WeightedRoundRobinAlgorithm{
		repo:  repo,
		index: 0,
	}
}

func (a *WeightedRoundRobinAlgorithm) ReverseProxy(w http.ResponseWriter, r *http.Request) {
	nextBackend := a.Next()
	// header := w.Header().Clone()
	log.Printf("[Weighted Round Robin] Next Backend ID: %v", nextBackend)
}

func (a *WeightedRoundRobinAlgorithm) Next() *entity.Backend {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.backends) == 0 {
		return nil
	}

	backend := a.backends[0]
	return &backend
}
