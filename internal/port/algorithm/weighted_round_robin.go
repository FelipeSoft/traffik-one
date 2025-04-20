package algorithm

import (
	"log"
	"net/http"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type WeightedRoundRobinAlgorithm struct {
	configEvent *entity.ConfigEvent
	index       uint32
	mu          sync.RWMutex
}

func NewWeightedRoundRobinAlgorithm(configEvent *entity.ConfigEvent) *WeightedRoundRobinAlgorithm {
	return &WeightedRoundRobinAlgorithm{
		configEvent: configEvent,
		index:       0,
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

	if len(a.configEvent.Backend) == 0 {
		return nil
	}

	backend := a.configEvent.Backend[0]
	return &backend
}
