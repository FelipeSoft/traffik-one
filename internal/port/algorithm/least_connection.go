package algorithm

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type LeastConnectionAlgorithm struct {
	repo     port.BackendRepository
	backends []entity.Backend
	index    uint32
	mu       sync.RWMutex
}

func NewLeastConnectionAlgorithm(repo port.BackendRepository) *LeastConnectionAlgorithm {
	return &LeastConnectionAlgorithm{
		repo:  repo,
		index: 0,
	}
}

func (a *LeastConnectionAlgorithm) ReverseProxy(w http.ResponseWriter, r *http.Request) {
	nextBackend := a.Next()
	// header := w.Header().Clone()
	log.Printf("[Classic Round Robin] Next Backend ID: %v", nextBackend)
}

func (a *LeastConnectionAlgorithm) Next() *entity.Backend {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.backends) == 0 {
		return nil
	}

	idx := atomic.AddUint32(&a.index, 1) - 1
	backend := a.backends[idx%uint32(len(a.backends))]
	return &backend
}
