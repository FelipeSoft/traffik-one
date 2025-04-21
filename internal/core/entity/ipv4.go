package entity

import (
	"fmt"
	"net"
)

type IPV4 struct {
	ipv4 string
}

func NewIPV4(ipv4 string) (*IPV4, error) {
	inputIPv4 := &IPV4{ipv4: ipv4}
	if inputIPv4.isIPv4() {
		return inputIPv4, nil
	}
	return nil, fmt.Errorf("invalid ipv4 '%s' provided", inputIPv4)
}

func (i *IPV4) isIPv4() bool {
	ip := net.ParseIP(i.ipv4)
	if ip == nil {
		return false
	}
	return ip.To4() != nil
}
