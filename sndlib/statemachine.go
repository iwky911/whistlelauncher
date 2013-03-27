package sndlib

// A simple state machine used to add inertia to the note detector and avoid detections errors
type State struct {
	limit   int
	counter int
}

func NewState(l int) *State {
	return &State{l, 0}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (s *State) Incr() {
	s.counter = min(s.counter+1, s.limit)
}

func (s *State) Decr() {
	s.counter = max(s.counter-1, -s.limit)
}

func (s *State) IsActive() bool {
	return s.counter > 0
}
