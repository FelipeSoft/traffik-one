package dispatcher

import (
	"context"
	"log"
	"sync"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type Dispatcher struct {
	configEvent *entity.ConfigEvent
	dispatchers []port.Dispatcher
	wg          sync.WaitGroup
}

func NewDispatcher(configEvent *entity.ConfigEvent, dispatchers []port.Dispatcher) *Dispatcher {
	return &Dispatcher{
		configEvent: configEvent,
		dispatchers: dispatchers,
	}
}

func (d *Dispatcher) Start(ctx context.Context) {
	for _, dispatcher := range d.dispatchers {
		d.wg.Add(1)
		go func(dispatcher port.Dispatcher) {
			defer d.wg.Done()
			dispatcher.Start(ctx)
		}(dispatcher)
	}

	go func() {
		<-ctx.Done()
		d.wg.Wait()
		log.Println("All dispatchers exited.")
	}()
}
