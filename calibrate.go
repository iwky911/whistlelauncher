package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/iwky911/whistlelauncher/persistance"
	"github.com/iwky911/whistlelauncher/sndlib"
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

func isPrefix(a, b string) bool {
	return len(a) >= len(b) && a[:len(b)] == b
}

/*
Acquire list of commands and whistle tones
*/
func configure() {
	name := ""
	command := ""
	var current_seq *sndlib.Sequence

	commandes := persistance.LoadFromFile(DATA_FILE)
	notechan := make(chan *sndlib.Note, 1)
	seqchan := make(chan *sndlib.Sequence, 1)
	cmdchan := make(chan string, 1)
	go sndlib.DetectNote(notechan)
	go sndlib.DetectSequence(notechan, seqchan)
	go acquireCommand(cmdchan)
	fmt.Println("Configuration tool. Type help for commands")
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
			case isPrefix(cmd, "name="):
				name = cmd[5:]
			case isPrefix(cmd, "cmd="):
				command = cmd[4:]
			case cmd == "confirm":
				if name =="" || command=="" {
					fmt.Println("name or command null")
					break
				}
				var list = []float64{}
				for _, n := range current_seq.Notes() {
					if n != nil {
						list = append(list, n.Value())
					}
				}
				commandes = append(commandes, persistance.Mapping{list, name, command})
			case cmd=="remove":
				newcommandes := []persistance.Mapping{}
				for _,c := range commandes {
					if c.Name != name{
						newcommandes = append(newcommandes, c)
					}
				}
				commandes=newcommandes
				fmt.Println(commandes)
			case cmd == "print":
				fmt.Println(commandes)

			case cmd == "save":
				persistance.SaveToFile(DATA_FILE, commandes)
			case cmd=="help":
				fmt.Println(`
List of the commands:
name=### (set the name of the command)
cmd=### (set the command that will be mapped)
remove (remove the commands whose name is the current name)
confirm=### (add the curent mapping (name, command and sound sequence) to the config)
save (save the config to the config file)
print (print the current config)
exit or quit (quit the program)
					`)
			default:
				fmt.Println("unknown command")
			}

		}
	}
}
