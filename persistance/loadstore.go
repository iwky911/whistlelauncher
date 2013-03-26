package persistance

import(
	"fmt"
	"encoding/gob"
	"os"
)
var _ = fmt.Print

func SaveToFile(name string, data []Mapping){
	f,err := os.OpenFile(name, os.O_WRONLY + os.O_CREATE, 0666)
	if err!= nil {
		fmt.Println(err)
	}
	defer f.Close()
	gob.NewEncoder(f).Encode(data)
}

func LoadFromFile(name string) []Mapping {
	f,err := os.Open(name)
	if err!= nil {
		fmt.Println(err)
		return []Mapping{}
	}
	defer f.Close()
	output := []Mapping{}
	gob.NewDecoder(f).Decode(&output)
	return output
}

