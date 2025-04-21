package algorithm

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type WeightedRoundRobinAlgorithm struct {
	configEvent *entity.ConfigEvent
	randGen     *rand.Rand
	randLock    sync.Mutex
}

func NewWeightedRoundRobinAlgorithm(configEvent *entity.ConfigEvent) *WeightedRoundRobinAlgorithm {
	return &WeightedRoundRobinAlgorithm{
		configEvent: configEvent,
		randGen:     rand.New(rand.NewSource(rand.Int63())),
	}
}

func (a *WeightedRoundRobinAlgorithm) ReverseProxy(w http.ResponseWriter, r *http.Request) {
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

func (a *WeightedRoundRobinAlgorithm) Next() *entity.Backend {
	var total int
	var activeBackends []*entity.Backend

	// Primeira passada: calcular peso total e coletar backends ativos
	for i := range a.configEvent.Backend {
		if a.configEvent.Backend[i].State {
			total += a.configEvent.Backend[i].Weight
			activeBackends = append(activeBackends, &a.configEvent.Backend[i])
		}
	}

	if total == 0 {
		return nil
	}

	// Gerar número aleatório seguro para concorrência
	a.randLock.Lock()
	randomValue := a.randGen.Intn(total)
	a.randLock.Unlock()

	// Segunda passada: encontrar o backend correspondente
	var cumulative int
	for _, backend := range activeBackends {
		cumulative += backend.Weight
		if randomValue < cumulative {
			return backend
		}
	}

	return nil // Nunca deve chegar aqui se total > 0
}