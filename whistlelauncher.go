package main

import (
	"fmt"
	"github.com/iwky911/whistlelauncher/persistance"
	"github.com/iwky911/whistlelauncher/sndlib"
)

func watchWhistle() {
	commandes := persistance.LoadFromFile(DATA_FILE)
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
