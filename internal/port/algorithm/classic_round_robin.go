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
	if nextBackend == nil {
		http.Error(w, "no available backends", http.StatusServiceUnavailable)
		return
	}

	var backendURL string

	backendSource := r.RequestURI
	if nextBackend.Hostname != "none" {
		backendURL = fmt.Sprintf("%s://%s%s", nextBackend.Protocol, nextBackend.Hostname, backendSource)
	} else {
		backendURL = fmt.Sprintf("%s://%s:%d%s", nextBackend.Protocol, nextBackend.IPv4, nextBackend.Port, backendSource)
	}

	routingRules := a.configEvent.RoutingRules
	if len(routingRules) > 0 {
		for _, rule := range routingRules {
			if rule.Source != backendSource {
				http.Error(w, fmt.Sprintf("error forwarding request to backend %s: source '%s' not found", nextBackend.ID, backendSource), http.StatusNotFound)
				return
			}
		}
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

func (a *ClassicRoundRobinAlgorithm) Next() *entity.Backend {
	if len(a.configEvent.Backend) == 0 {
		return nil
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	idx := atomic.AddUint32(&a.index, 1)

	if idx >= uint32(len(a.configEvent.Backend)) {
		atomic.StoreUint32(&a.index, 0)
		idx = 0
	}

	backend := a.configEvent.Backend[idx%uint32(len(a.configEvent.Backend))]
	return &backend
}
