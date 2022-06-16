package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type DecisiontreeIntruction struct {
	Name           string `json:"name"`
	Kind           string `json:"kind"`
	Criterion      string `json:"criterion"`
	MinLeafSamples int    `json:"minLeafSamples"`
}

func ParseInstruction1(f multipart.File) []DecisiontreeIntruction {
	filebytes, _ := ioutil.ReadAll(f)
	var instructions []DecisiontreeIntruction
	err := json.Unmarshal(filebytes, &instructions)
	if err != nil {
		log.Fatal(err)
	}

	return instructions
}
