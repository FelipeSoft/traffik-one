package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port/websocket"
)

func StartHttpHealthChecker(ctx context.Context, ws *websocket.WebsocketServer, configEvent *entity.ConfigEvent, interval time.Duration, workers int) {
	var wg sync.WaitGroup
	workChan := make(chan *entity.Backend, workers)

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for backend := range workChan {
				log.Printf("Backend: %v", backend)
				marshalBackend, err := json.Marshal(backend)
				if err != nil {
					log.Printf("Error to marshal backend: %v", err)
				}
				ws.Broadcast(marshalBackend)
				checkBackendHealth(backend)
			}
		}()
	}

	ticker := time.NewTicker(interval)
	go func() {
		defer close(workChan)
		for {
			select {
			case <-ticker.C:
				configEvent.Mu.Lock()
				for i := range configEvent.Backend {
					select {
					case workChan <- &configEvent.Backend[i]:
					case <-ctx.Done():
						return
					}
				}
				configEvent.Mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
}

func checkBackendHealth(b *entity.Backend) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	healthURL := fmt.Sprintf("%s://%s:%d/health", b.Protocol, b.Hostname, b.Port)
	if b.Hostname == "none" {
		healthURL = fmt.Sprintf("%s://%s:%d/health", b.Protocol, b.IPv4, b.Port)
	}

	resp, err := client.Get(healthURL)
	if err != nil {
		// log.Printf("Health check failed for backend %s: %v", b.ID, err)
		if err := b.Inactivate(); err != nil {
			// log.Printf("Failed to deactivate backend %s: %v", b.ID, err)
		}
		return
	}
	defer resp.Body.Close()

    activeBackend := b.State
    inactiveBackend := !activeBackend

	if resp.StatusCode == http.StatusOK {
		if activeBackend {
            return
        }
        b.Activate()
	}

	if resp.StatusCode != http.StatusOK {
		if inactiveBackend {
            return
        }
        b.Inactivate()
	}
}