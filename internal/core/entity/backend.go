package entity

import (
	"fmt"
	"log"

	"github.com/FelipeSoft/traffik-one/internal/port/idgen"
)

type Backend struct {
	ID       string `json:"id"`
	IPv4     string `json:"ipv4"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Weight   int    `json:"weight"`
	State    bool   `json:"state"`
	PoolID   string `json:"poolId"`
}

func NewBackend(ipv4 string, hostname string, port int, protocol string, weight int, poolId string) (*Backend, error) {
	validIPv4, err := NewIPV4(ipv4)
	if err != nil {
		return nil, err
	}

	validWeight, err := NewWeight(weight)
	if err != nil {
		return nil, err
	}

	validPort, err := NewPort(port)
	if err != nil {
		return nil, err
	}

	validProtocol, err := NewProtocol(protocol)
	if err != nil {
		return nil, err
	}

	validHostname, err := NewHostname(hostname)
	if err != nil {
		return nil, err
	}

	id := idgen.GenerateID()
	backend := &Backend{
		ID:       id,
		IPv4:     validIPv4.ipv4,
		Hostname: validHostname.hostname,
		Port:     validPort.port,
		Protocol: validProtocol.protocol,
		Weight:   validWeight.weight,
		PoolID:   poolId,
	}
	backend.Activate()
	return backend, nil
}

func (b *Backend) Update(ipv4 string, hostname string, port int, protocol string, weight int, poolId string) (*Backend, error) {
	backend := b
	if b.IPv4 != ipv4 && ipv4 != "" {
		validIPv4, err := NewIPV4(ipv4)
		if err != nil {
			return nil, err
		}
		backend.IPv4 = validIPv4.ipv4
	}
	if b.Weight != weight && weight != 0 {
		validWeight, err := NewWeight(weight)
		if err != nil {
			return nil, err
		}
		backend.Weight = validWeight.weight
	}
	if b.Port != port && port != 0 {
		validPort, err := NewPort(port)
		if err != nil {
			return nil, err
		}
		backend.Port = validPort.port
	}
	if b.Protocol != protocol && protocol != "" {
		validProtocol, err := NewProtocol(protocol)
		if err != nil {
			return nil, err
		}
		backend.Protocol = validProtocol.protocol
	}
	if b.Hostname != hostname && hostname != "" {
		validHostname, err := NewHostname(hostname)
		if err != nil {
			return nil, err
		}
		backend.Hostname = validHostname.hostname
	}
	if poolId == "0" {
		backend.PoolID = b.PoolID
	}
	log.Print(backend)
	return backend, nil
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
