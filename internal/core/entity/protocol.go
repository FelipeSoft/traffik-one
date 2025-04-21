package entity

import "slices"

import "fmt"

type Protocol struct {
	protocol string
}

var availableProtocols = []string{"http", "https"}

func NewProtocol(protocol string) (*Protocol, error) {
	valid := slices.Contains(availableProtocols, protocol)
	if valid {
		return &Protocol{protocol: protocol}, nil
	}
	return nil, fmt.Errorf("invalid protocol '%s' provided for backend", protocol)
}