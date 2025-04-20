package dispatcher

import (
	"context"
	"log"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type BackendDispatcher struct {
	configEvent *entity.ConfigEvent
	ch          chan []entity.Backend
	mu          sync.Mutex
	closed      bool
	muClose     sync.Mutex
}

func NewBackendDispatcher(configEvent *entity.ConfigEvent, ch chan []entity.Backend) *BackendDispatcher {
	return &BackendDispatcher{
		configEvent: configEvent,
		ch:          ch,
	}
}

func (d *BackendDispatcher) Start(ctx context.Context) {
	defer func() {
		d.muClose.Lock()
		if !d.closed {
			close(d.ch)
			d.closed = true
		}
		d.muClose.Unlock()
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("BackendDispatcher finished by context cancelling.")
			return
		case backend, ok := <-d.ch:
			if !ok {
				log.Println("Backends dispatcher channel closed. Finishing BackendDispatcher.")
				return
			}
			d.mu.Lock()
			d.configEvent.Backend = backend
			d.mu.Unlock()
		}
	}
}

func (d *BackendDispatcher) Dispatch(args any) {
	backends, ok := args.([]entity.Backend)
	if !ok {
		log.Fatalf("could not proceed with invalid dispatch assertion for backend dispatcher")
	}

	d.muClose.Lock()
	defer d.muClose.Unlock()
	if d.closed {
		log.Println("backend channel is already closed, skipping dispatch")
		return
	}

	select {
	case d.ch <- backends:
	default:
		log.Println("Backend dispatcher channel full or not ready")
	}
}
