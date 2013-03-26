package persistance

import(
	"os/exec"
	"strconv"
	. "github.com/iwky911/whistlelauncher/tools"
)

type Mapping struct{
	Notes []float64
	Name, Command string
}

func (m *Mapping) Execute(){
	D(strconv.Quote(m.Command))
	cmd := exec.Command("/bin/sh","-c", "\""+strconv.Quote(m.Command)+"\"")
	cmd.Start()
}