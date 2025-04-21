package entity

import "fmt"

type Port struct {
	port int
}

func NewPort(port int) (*Port, error) {
	if port >= 0 && port <= 65535 {
		return nil, fmt.Errorf("invalid port '%d' provided for backend", port)
	}
	return &Port{port: port}, nil
}