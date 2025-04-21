package entity

import "fmt"

type Hostname struct {
	hostname string
}

func NewHostname(hostname string) (*Hostname, error) {
	if hostname == "" {
		return nil, fmt.Errorf("empty hostname provided")
	}
	if len(hostname) > 255 {
		return nil, fmt.Errorf("invalid hostname '%s' provided", hostname)
	}
	return &Hostname{hostname: hostname}, nil
}
