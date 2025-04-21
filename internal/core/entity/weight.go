package entity

import (
	"fmt"
)

type Weight struct {
	weight int
}

func NewWeight(weight int) (*Weight, error) {
	if weight < 1 || weight > 100 {
		return nil, fmt.Errorf("the weight of the backend must be in the range of 1 and 100")
	}
	return &Weight{weight: weight}, nil
}
