package dispatcher

import (
	"context"
	"log"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type AlgorithmsDispatcher struct {
	configEvent *entity.ConfigEvent
	ch          chan string
	mu          sync.Mutex
	closed      bool
	muClose     sync.Mutex
}

func NewAlgorithmsDispatcher(configEvent *entity.ConfigEvent, ch chan string) *AlgorithmsDispatcher {
	return &AlgorithmsDispatcher{
		configEvent: configEvent,
		ch:          ch,
	}
}

func (d *AlgorithmsDispatcher) Start(ctx context.Context) {
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
			log.Println("AlgorithmDispatcher finished by context cancelling.")
			return
		case algorithm, ok := <-d.ch:
			if !ok {
				log.Println("Algorithm dispatcher channel closed. Finishing AlgorithmDispatcher.")
				return
			}
			d.mu.Lock()
			d.configEvent.Algorithm = &algorithm
			d.mu.Unlock()
		}
	}
}

func (d *AlgorithmsDispatcher) Dispatch(args any) {
	algorithm, ok := args.(string)
	if !ok {
		log.Fatalf("could not proceed with invalid dispatch assertion for algorithm dispatcher")
	}

	d.muClose.Lock()
	defer d.muClose.Unlock()
	if d.closed {
		log.Println("algorithm channel is already closed, skipping dispatch")
		return
	}

	select {
	case d.ch <- algorithm:
	default:
		log.Println("Algorithm dispatcher channel full or not ready")
	}
}
