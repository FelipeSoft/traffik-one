package entity

import (
	"time"

	"github.com/FelipeSoft/traffik-one/internal/port/idgen"
)

type RoutingRules struct {
	ID        string    `json:"id"`
	Source    string    `json:"source"`
	Target    string    `json:"target"`
	Protocol  string    `json:"protocol"`
	PoolID    string    `json:"poolId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewRoutingRules(source string, target string, protocol string, poolID string) *RoutingRules {
	id := idgen.GenerateID()
	httpRoutingRule := &RoutingRules{
		ID:        id,
		Source:    source,
		Target:    target,
		Protocol:  protocol,
		PoolID:    poolID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return httpRoutingRule
}
