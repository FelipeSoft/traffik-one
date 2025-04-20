package algorithm

import (
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type AlgorithmFactory struct {
	configEvent *entity.ConfigEvent
}

func NewAlgorithmFactory(configEvent *entity.ConfigEvent) *AlgorithmFactory {
	return &AlgorithmFactory{
		configEvent: configEvent,
	}
}

func (f *AlgorithmFactory) Create() (port.Algorithm, error) {
	algorithm := *f.configEvent.Algorithm
	switch algorithm {
	case "crr":
		return NewClassicRoundRobinAlgorithm(f.configEvent), nil
	case "wrr":
		return NewWeightedRoundRobinAlgorithm(f.configEvent), nil
	case "lc0":
		return NewLeastConnectionAlgorithm(f.configEvent), nil
	default:
		return nil, fmt.Errorf("invalid algorithm '%s'", algorithm)
	}
}
