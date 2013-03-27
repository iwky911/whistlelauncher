package sndlib

import (
	"time"
)

var NB_STATE = 2 // nombre of state for the state machine

func DetectNote(output chan *Note) {
	sp := CreateInstance()
	state := NewState(NB_STATE)
	var current_note *Note
	for true {
		s := sp.GetSndSignature()
		if s.IsClear() {
			state.Incr()
		} else {
			state.Decr()
		}

		if state.IsActive() {
			if current_note == nil {
				current_note = &Note{s.Value(), time.Now(), time.Now()}
			}
		} else {
			if current_note != nil {
				current_note.end = time.Now()
				output <- current_note
			}
			current_note = nil
		}
	}
}
