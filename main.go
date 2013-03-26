package main

import(
	"os"
)

var DATA_FILE = "mapping.cfg"

func main(){
	var config = false
	for _,arg := range os.Args {
		if arg == "config" {
			config = true
		}
	}
	if config {
		configure()
	}else{
		watchWhistle()
	}
}