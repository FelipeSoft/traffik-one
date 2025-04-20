package algorithm

import (
	"log"

	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type Container struct {
	ClassicRoundRobinAlgorithm  port.Algorithm
	WeightedRoundRobinAlgorithm port.Algorithm
	LeastConnectionAlgorithm    port.Algorithm
}

func NewContainer() *Container {
	classicRoundRobinAlgorithm := NewClassicRoundRobinAlgorithm()
	weightedRoundRobinAlgorithm := NewWeightedRoundRobinAlgorithm()
	leastConnectionAlgorithm := NewLeastConnectionAlgorithm()

	log.Println("[Algorithms Container] Dependencies loaded successfully")

	return &Container{
		ClassicRoundRobinAlgorithm:  classicRoundRobinAlgorithm,
		WeightedRoundRobinAlgorithm: weightedRoundRobinAlgorithm,
		LeastConnectionAlgorithm:    leastConnectionAlgorithm,
	}
}
