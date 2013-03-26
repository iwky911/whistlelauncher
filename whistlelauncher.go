package main

import (
	"fmt"
	"persistance"
	"sndlib"
)

func main() {
	commandes := persistance.LoadFromFile("mapping.cfg")
	mychan := make(chan *sndlib.Note, 1)
	seqchan := make(chan *sndlib.Sequence, 1)
	go sndlib.DetectNote(mychan)
	go sndlib.DetectSequence(mychan, seqchan)

	for {
		seq := <-seqchan
		for _, c := range commandes {
			if seq.Matches(c.Notes) {
				fmt.Println("commande recognized: " + c.Name)
				break
			}
		}
		fmt.Println("Sequence " + seq.String())
	}
}
