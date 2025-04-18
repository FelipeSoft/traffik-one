package entity

import (
	"time"

	"github.com/FelipeSoft/traffik-one/internal/port/idgen"
)

type RoutingRules struct {
	ID        string
	Source    string
	Target    string
	Protocol  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRoutingRules(source string, target string, protocol string) *RoutingRules {
	id := idgen.GenerateID()
	httpRoutingRule := &RoutingRules{
		ID:        id,
		Source:    source,
		Target:    target,
		Protocol:  protocol,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return httpRoutingRule
}
