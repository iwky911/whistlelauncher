package main

import(
	"os"
)

// Filepath for the config file
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