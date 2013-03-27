package sndlib

import (
	"time"
)

var MIN_NOTE_LENGTH = .05 // Length of the second note
var MIN_START_NOTE_LENGTH = .05 // Length of the first note
var MAX_TIMEOUT = 750*time.Millisecond // Max delay between the end of two notes


//Detect sequence of notes out of a note stream
func DetectSequence(input chan *Note, output chan *Sequence) {
	var sequence = &Sequence{[3]*Note{nil, nil, nil}, 0}
	var timeout = make(chan *Note)
	for true {
		send := false
		select {
		case n := <-input:
			if n.Length() > MIN_START_NOTE_LENGTH || sequence.IsEmpty() && n.Length() > MIN_NOTE_LENGTH {
				sequence.addNote(n)
				time.AfterFunc(MAX_TIMEOUT, func() {
					timeout <- n
				})
			}
		case n := <-timeout:
			if sequence.isLast(n) && !sequence.IsEmpty() {
				send = true
			}
		}
		if sequence.isFull() || send == true {
			output <- sequence
			sequence = &Sequence{[3]*Note{}, 0}
			send = false
		}
	}
}
