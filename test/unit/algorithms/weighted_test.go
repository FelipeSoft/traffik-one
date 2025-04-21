package algorithms_test

import (
	"math"
	"testing"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/port/algorithm"
)

func TestWeightedDistribution(t *testing.T) {
	backends := []entity.Backend{
		{Weight: 30, State: true},
		{Weight: 70, State: true},
	}

	config := &entity.ConfigEvent{Backend: backends}
	algo := algorithm.NewWeightedRoundRobinAlgorithm(config)

	const totalRequests = 1_000_000
	results := make(map[*entity.Backend]int)

	for range totalRequests {
		backend := algo.Next()
		if backend == nil {
			t.Fatal("Next() returned nil but backends are active")
		}
		results[backend]++
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 backends, got %d", len(results))
	}

	const tolerance = 0.01
	for backend, count := range results {
		expected := float64(backend.Weight) / 100
		actual := float64(count) / totalRequests
		diff := math.Abs(expected - actual)

		if diff > tolerance {
			t.Errorf(
				"Backend %s (weight %d): expected %.2f%%, got %.2f%% (difference %.2f%%)",
				backend.ID,
				backend.Weight,
				expected*100,
				actual*100,
				diff*100,
			)
		}
	}
}