package entity

import (
	"fmt"

	"github.com/FelipeSoft/traffik-one/internal/port/idgen"
)

type Backend struct {
	ID       string  `json:"id"`
	IPv4     string  `json:"ipv4"`
	Hostname *string `json:"hostname"`
	Port     int     `json:"port"`
	Protocol string  `json:"protocol"`
	Weight   *int    `json:"weight"`
	State    bool    `json:"state"`
	PoolID   string `json:"poolId"`
}

func NewBackend(ipv4 string, hostname *string, port int, protocol string, weight *int, poolId string) *Backend {
	id := idgen.GenerateID()
	backend := &Backend{
		ID:       id,
		IPv4:     ipv4,
		Hostname: hostname,
		Port:     port,
		Protocol: protocol,
		Weight:   weight,
		PoolID:   poolId,
	}
	backend.Activate()
	return backend
}

func (b *Backend) AssignToPool(poolId string) {
	b.PoolID = poolId
}

func (b *Backend) Activate() error {
	if b.State {
		return fmt.Errorf("[core error] the backend is already activated")
	}
	b.State = true
	return nil
}

func (b *Backend) Inactivate() error {
	if !b.State {
		return fmt.Errorf("[core error] the backend is already inactivated")
	}
	b.State = false
	return nil
}
