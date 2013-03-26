package sndlib

import (
	"encoding/gob"
	"fmt"
	"math"
	"os"
)

type Samples interface {
	GetCloser([]float64) string
}

type mysamples struct {
	samps []sample
}

type sample struct {
	name  string
	value []float64
}

func LoadSamples() Samples {
	dir, err := os.Open("data")
	if err != nil {
		fmt.Println("error")
	}
	samples := []sample{}
	defer dir.Close()
	filenames, _ := dir.Readdirnames(-1)
	for _, filename := range filenames {
		f, _ := os.Open("data/" + filename)
		defer f.Close()

		decoder := gob.NewDecoder(f)
		value := make([]float64, 0, 0)
		decoder.Decode(&value)
		if filename != "void" {
			fmt.Println(filename)
			samples = append(samples, sample{filename, value})
		}
	}
	return mysamples{samples}
}

func (s sample) Diff(other []float64) float64 {
	a := 0.
	for i, _ := range other {
		a += math.Pow(other[i]-s.value[i], 2)
	}
	return math.Sqrt(a)
}
func (s sample) GetName() string {
	return s.name
}

func (s mysamples) GetCloser(v []float64) string {
	rec := s.samps[0].Diff(v)
	name := s.samps[0].GetName()
	for i, _ := range s.samps {
		d := s.samps[i].Diff(v)
		if d < rec {
			rec = d
			name = s.samps[i].GetName()
		}
	}
	return name
}
