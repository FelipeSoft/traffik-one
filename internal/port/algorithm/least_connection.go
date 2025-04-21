package algorithm

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
	if nextBackend == nil {
		http.Error(w, "no available backends", http.StatusServiceUnavailable)
		return
	}

	var backendURL string
	if nextBackend.Hostname != "none" {
		backendURL = fmt.Sprintf("%s://%s%s", nextBackend.Protocol, nextBackend.Hostname, r.RequestURI)
	} else {
		backendURL = fmt.Sprintf("%s://%s:%d%s", nextBackend.Protocol, nextBackend.IPv4, nextBackend.Port, r.RequestURI)
	}

	req, err := http.NewRequest(r.Method, backendURL, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("backend %s error: %v", nextBackend.ID, err), http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		req.Header.Set(key, strings.Join(values, ", "))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error forwarding request to backend %s: %v", nextBackend.ID, err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		w.Header().Set(key, strings.Join(values, ", "))
	}
	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("error copying response body: %v", err)
	}
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
