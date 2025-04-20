package entity

type ConfigEvent struct {
	RoutingRules []RoutingRules
	Backend      []Backend
	Algorithm    *string // wrr, crr, lc0
}
