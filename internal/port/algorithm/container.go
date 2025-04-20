package algorithm

import (
	"log"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type Container struct {
	ClassicRoundRobinAlgorithm  port.Algorithm
	WeightedRoundRobinAlgorithm port.Algorithm
	LeastConnectionAlgorithm    port.Algorithm
}

func NewContainer(configEvent *entity.ConfigEvent) *Container {
	classicRoundRobinAlgorithm := NewClassicRoundRobinAlgorithm(configEvent)
	weightedRoundRobinAlgorithm := NewWeightedRoundRobinAlgorithm(configEvent)
	leastConnectionAlgorithm := NewLeastConnectionAlgorithm(configEvent)

	log.Println("[Algorithms Container] Dependencies loaded successfully")

	return &Container{
		ClassicRoundRobinAlgorithm:  classicRoundRobinAlgorithm,
		WeightedRoundRobinAlgorithm: weightedRoundRobinAlgorithm,
		LeastConnectionAlgorithm:    leastConnectionAlgorithm,
	}
}
