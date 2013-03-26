package sndlib

import (
	"fmt"
	"math"
)

const NOTE_MATCHING_THRESHOLD = 3

type Sequence struct {
	notes [3]*Note
	i     int
}

func (s *Sequence) addNote(n *Note) {
	//fmt.Println("debug: note added: ")
	s.notes[s.i] = n
	s.i++
}

func (s *Sequence) isLast(n *Note) bool {
	return s.i > 0 && s.notes[s.i-1] == n
}

func (s *Sequence) flush() {
	for i, _ := range s.notes {
		s.notes[i] = nil
	}
	s.i = 0
}

func (s *Sequence) isFull() bool {
	return s.i >= len(s.notes)
}

func (s *Sequence) IsEmpty() bool {
	return s.i == 0
}

func (s *Sequence) Notes() []*Note {
	return s.notes[:]
}

func (s *Sequence) Matches(notes []float64) bool {
	for i, n := range s.notes {
		if n != nil && (len(notes) <= i || math.Abs(n.Value()-notes[i]) > NOTE_MATCHING_THRESHOLD) {
			return false
		}
		if n == nil && len(notes) > i {
			return false
		}
	}
	return true
}

func (s *Sequence) String() string {
	sortie := ""
	for _, v := range s.notes {
		if v != nil {
			sortie += fmt.Sprintf("%.2f ", v.Value())
		}
	}
	return sortie
}
