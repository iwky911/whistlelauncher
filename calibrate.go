package main

import (
	"bufio"
	"fmt"
	"os"
	"persistance"
	"sndlib"
)

func printSeq(s *sndlib.Sequence, name string) {
	fmt.Print("sequence for " + name + ": ")
	for _, v := range s.Notes() {
		if v != nil {
			fmt.Printf("%f ", v.Value())
		}
	}
	fmt.Println()
}

func acquireCommand(cmdchan chan string) {
	r := bufio.NewReader(os.Stdin)
	for true {
		cmd, _ := r.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		cmdchan <- cmd

		if cmd == "exit" || cmd == "quit" {
			break
		}
	}
}

func isPrefix(a, b string, l int) bool {
	return len(a) >= l && a[:l] == b
}

func main() {
	name := ""
	command := ""
	var current_seq *sndlib.Sequence

	commandes := []persistance.Mapping{}
	notechan := make(chan *sndlib.Note, 1)
	seqchan := make(chan *sndlib.Sequence, 1)
	cmdchan := make(chan string, 1)
	go sndlib.DetectNote(notechan)
	go sndlib.DetectSequence(notechan, seqchan)
	go acquireCommand(cmdchan)
	for {
		fmt.Print(name + " -> " + command + "> ")
		select {
		case s := <-seqchan:
			current_seq = s
			printSeq(s, name)
		case cmd := <-cmdchan:
			switch {
			case cmd == "exit" || cmd == "quit":
				os.Exit(0)
			case isPrefix(cmd, "set name ", 9):
				name = cmd[9:]
			case isPrefix(cmd, "set cmd ", 8):
				command = cmd[8:]
			case cmd == "confirm":
				var list = []float64{}
				for _, n := range current_seq.Notes() {
					if n != nil {
						list = append(list, n.Value())
					}
				}
				commandes = append(commandes, persistance.Mapping{list, name, command})
			case cmd == "print":
				fmt.Println(commandes)
			case cmd == "save":
				persistance.SaveToFile("mapping.cfg", commandes)
			default:
				fmt.Println("unknown command")
			}

		}
	}
}
