package algorithm

import (
	"log"

	"github.com/FelipeSoft/traffik-one/internal/adapter/repository"
	"github.com/FelipeSoft/traffik-one/internal/core/port"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
)

type Container struct {
	ClassicRoundRobinAlgorithm  port.Algorithm
	WeightedRoundRobinAlgorithm port.Algorithm
	LeastConnectionAlgorithm    port.Algorithm
}

func NewContainer() *Container {
	boltDB := bolt.DB()
	backendRepository := repository.NewBoltBackendRepository(boltDB)

	classicRoundRobinAlgorithm := NewClassicRoundRobinAlgorithm(backendRepository)
	weightedRoundRobinAlgorithm := NewWeightedRoundRobinAlgorithm(backendRepository)
	leastConnectionAlgorithm := NewLeastConnectionAlgorithm(backendRepository)

	log.Println("[Algorithms Container] Dependencies loaded successfully")

	return &Container{
		ClassicRoundRobinAlgorithm:  classicRoundRobinAlgorithm,
		WeightedRoundRobinAlgorithm: weightedRoundRobinAlgorithm,
		LeastConnectionAlgorithm:    leastConnectionAlgorithm,
	}
}
