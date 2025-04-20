package algorithm

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type ClassicRoundRobinAlgorithm struct {
	configEvent *entity.ConfigEvent
	index       uint32
	mu          sync.RWMutex
}

func NewClassicRoundRobinAlgorithm(configEvent *entity.ConfigEvent) *ClassicRoundRobinAlgorithm {
	return &ClassicRoundRobinAlgorithm{
		configEvent: configEvent,
		index:       0,
	}
}

func (a *ClassicRoundRobinAlgorithm) ReverseProxy(w http.ResponseWriter, r *http.Request) {
	nextBackend := a.Next()
	// header := w.Header().Clone()
	log.Printf("[Classic Round Robin] Next Backend ID: %v", nextBackend)
}

func (a *ClassicRoundRobinAlgorithm) Next() *entity.Backend {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.configEvent.Backend) == 0 {
		return nil
	}

	idx := atomic.AddUint32(&a.index, 1) - 1
	backend := a.configEvent.Backend[idx%uint32(len(a.configEvent.Backend))]
	return &backend
}
