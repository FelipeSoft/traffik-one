package entity

import "fmt"

type Protocol struct {
	protocol string
}

var availableProtocols = []string{"http", "https"}

func NewProtocol(protocol string) (*Protocol, error) {
	valid := false
	for _, p := range availableProtocols {
		if protocol == p {
			valid = true
			break
		}
	}
	if valid {
		return &Protocol{protocol: protocol}, nil
	}
	return nil, fmt.Errorf("invalid protocol '%s' provided for backend", protocol)
}