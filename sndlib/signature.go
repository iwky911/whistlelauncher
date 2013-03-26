package sndlib

import (
	"math"
)

var CLEAR_THRESHOLD = 20.

type Signature interface {
	IsClear() bool
	Value() float64
}

type simpleSign struct {
	values [2]float64
}

func (s *simpleSign) IsClear() bool {
	return math.Abs(s.values[1]-s.values[0]) < CLEAR_THRESHOLD
}

func (s *simpleSign) Value() float64 {
	return (s.values[0] + s.values[1]) / 2
}
