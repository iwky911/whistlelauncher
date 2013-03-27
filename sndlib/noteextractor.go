/*
Package sndlib extracts notes or sequences of notes from sndpeek output.

	The three main components are:
		* CreateInstance that create the sndpeek instance and initialize the buffer
		* GetSndSignature that extract the two rolloffs from sndpeek output
		* DetectNote that extract notes from the sequence of rolloffs
		* DetectSequence that detect sequences of notes from sequence of notes

*/
package sndlib

import (
	"bufio"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"fmt"
	"os"
)


var whitespaces, _ = regexp.Compile("\\s+")

type SndPeekInstance struct {
	command *exec.Cmd
	reader  *bufio.Reader
}

func CreateInstance() *SndPeekInstance {
	c := exec.Command("sndpeek", "--nodisplay", "--print")

	pipe, err := c.StdoutPipe()
	if err != nil {
		fmt.Println("could not create pipe")
		fmt.Println(err)
		os.Exit(1)
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &SndPeekInstance{c, bufio.NewReader(pipe)}
}
// Returns a signature (two rolloffs)
func (sndpeek *SndPeekInstance) GetSndSignature() Signature {
	line, _ := sndpeek.reader.ReadString('\n')
	line = whitespaces.ReplaceAllString(line, " ")
	terms := strings.Split(line, " ")
	terms = terms[3:5]
	low, _ := strconv.ParseFloat(terms[0], 10)
	high, _ := strconv.ParseFloat(terms[1], 10)
	return &simpleSign{[2]float64{low, high}}
}
