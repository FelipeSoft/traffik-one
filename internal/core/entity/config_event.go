package entity

import "sync"

type ConfigEvent struct {
	RoutingRules []RoutingRules
	Backend      []Backend
	Algorithm    *string // wrr, crr, lc0
	Mu           sync.RWMutex
}
