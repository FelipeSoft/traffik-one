package algorithm

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type LeastConnectionAlgorithm struct {
	configEvent *entity.ConfigEvent
	index       uint32
	mu          sync.RWMutex
}

func NewLeastConnectionAlgorithm(configEvent *entity.ConfigEvent) *LeastConnectionAlgorithm {
	return &LeastConnectionAlgorithm{
		configEvent: configEvent,
		index:       0,
	}
}

func (a *LeastConnectionAlgorithm) ReverseProxy(w http.ResponseWriter, r *http.Request) {
	nextBackend := a.Next()
	// header := w.Header().Clone()
	log.Printf("[Least Connection] Next Backend ID: %v", nextBackend)
}

func (a *LeastConnectionAlgorithm) Next() *entity.Backend {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.configEvent.Backend) == 0 {
		return nil
	}

	idx := atomic.AddUint32(&a.index, 1) - 1
	backend := a.configEvent.Backend[idx%uint32(len(a.configEvent.Backend))]
	return &backend
}
