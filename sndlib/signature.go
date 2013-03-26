package sndlib

import (
	"fmt"
	"math"
)

const threshold = 4

type Signature interface {
	IsClear() bool
	Value() float64
	ToString() string
}

type simpleSign struct {
	values [2]float64
}

func (s *simpleSign) IsClear() bool {
	return math.Abs(s.values[1]-s.values[0]) < 30
}

// func (s *simpleSign) Matches(s2 Signature) bool {
// 	return math.Abs(s.Value() - s2.Value())< threshold
// }

func (s *simpleSign) Value() float64 {
	return (s.values[0] + s.values[1]) / 2
}

func (s *simpleSign) ToString() string {
	return fmt.Sprintf("%f", s.Value())
}
