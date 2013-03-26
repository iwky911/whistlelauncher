package sndlib

import (
	"fmt"
	"time"
)

const MIN_NOTE_LENGTH = .05
const MIN_START_NOTE_LENGTH = .05
const MAX_TIMEOUT = 500

var _ = fmt.Print

func DetectSequence(input chan *Note, output chan *Sequence) {
	var sequence = &Sequence{[3]*Note{nil, nil, nil}, 0}
	var timeout = make(chan *Note)
	for true {
		send := false
		select {
		case n := <-input:
			if n.Length() > MIN_START_NOTE_LENGTH || sequence.IsEmpty() && n.Length() > MIN_NOTE_LENGTH {
				sequence.addNote(n)
				time.AfterFunc(MAX_TIMEOUT*time.Millisecond, func() {
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
