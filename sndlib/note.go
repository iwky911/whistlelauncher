package sndlib

import (
	"time"
)

// A detected note
type Note struct {
	value float64 // tone
	start time.Time
	end   time.Time
}

func (n *Note) Length() float64 {
	return n.end.Sub(n.start).Seconds()
}

func (n *Note) Value() float64 {
	return n.value
}
