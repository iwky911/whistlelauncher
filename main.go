package main

import(
	"os"
)

var DATA_FILE = "data/mapping.cfg"

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