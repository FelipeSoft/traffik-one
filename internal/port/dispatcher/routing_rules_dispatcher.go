package dispatcher

import (
	"context"
	"log"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type RoutingRulesDispatcher struct {
	configEvent *entity.ConfigEvent
	ch          chan []entity.RoutingRules
	mu          sync.Mutex
	closed      bool
	muClose     sync.Mutex
}

func NewRoutingRulesDispatcher(configEvent *entity.ConfigEvent, ch chan []entity.RoutingRules) *RoutingRulesDispatcher {
	return &RoutingRulesDispatcher{
		configEvent: configEvent,
		ch:          ch,
	}
}

func (d *RoutingRulesDispatcher) Start(ctx context.Context) {
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
			log.Println("RoutingRulesDispatcher finished by context cancelling.")
			return
		case routingRules, ok := <-d.ch:
			if !ok {
				log.Println("Routing Rules dispatcher channel closed. Finishing RoutingRulesDispatcher.")
				return
			}
			d.mu.Lock()
			d.configEvent.RoutingRules = routingRules
			d.mu.Unlock()
		}
	}
}

func (d *RoutingRulesDispatcher) Dispatch(args any) {
	routingRules, ok := args.([]entity.RoutingRules)
	if !ok {
		log.Fatalf("could not proceed with invalid dispatch assertion for backend dispatcher")
	}

	d.muClose.Lock()
	defer d.muClose.Unlock()
	if d.closed {
		log.Println("routing rules channel is already closed, skipping dispatch")
		return
	}

	select {
	case d.ch <- routingRules:
	default:
		log.Println("Routing Rules dispatcher channel full or not ready")
	}
}
