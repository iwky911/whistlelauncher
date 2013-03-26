package tools

import(
	"os"
	"fmt"
)

var debug = os.Getenv("DEBUG")=="1"

func D(a ...interface{}) {
	fmt.Print("debug: ")
	fmt.Println(a)
}