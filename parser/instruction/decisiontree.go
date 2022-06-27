package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type DecisiontreeRegIntruction struct {
	Name           string `json:"name"`
	Kind           string `json:"kind"`
	Criterion      string `json:"criterion"`
	MinLeafSamples int    `json:"minLeafSamples"`
}

type DecisiontreeClassIntruction struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Criterion string `json:"criterion"`
}

func ParseDTRegInstruction(f multipart.File) []DecisiontreeRegIntruction {
	filebytes, _ := ioutil.ReadAll(f)
	var instructions []DecisiontreeRegIntruction
	err := json.Unmarshal(filebytes, &instructions)
	if err != nil {
		log.Fatal(err)
	}

	return instructions
}

func ParseDTClassInstruction(f multipart.File) []DecisiontreeClassIntruction {
	filebytes, _ := ioutil.ReadAll(f)
	var instructions []DecisiontreeClassIntruction
	err := json.Unmarshal(filebytes, &instructions)
	if err != nil {
		log.Fatal(err)
	}

	return instructions
}
