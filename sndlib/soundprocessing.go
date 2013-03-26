package sndlib

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const INIT_NB = 25

var whitespaces, _ = regexp.Compile("\\s+")
var meanvector []float64

func nothing() {
	fmt.Println("do nothing")
}

// func add(a, b []float64) {
// 	for i, _ := range a {
// 		a[i] += b[i]
// 	}
// }

// func average(v []float64, n float64) {
// 	for i, _ := range v {
// 		v[i] = v[i] / n
// 	}
// }

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
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err)
	}
	return &SndPeekInstance{c, bufio.NewReader(pipe)}
}

func (sndpeek *SndPeekInstance) GetSndSignature() Signature {
	line, _ := sndpeek.reader.ReadString('\n')
	line = whitespaces.ReplaceAllString(line, " ")
	terms := strings.Split(line, " ")
	terms = terms[3:5]
	low, _ := strconv.ParseFloat(terms[0], 10)
	high, _ := strconv.ParseFloat(terms[1], 10)
	return &simpleSign{[2]float64{low, high}}
}
