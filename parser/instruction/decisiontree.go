package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
)

// represents the training instructions for a regression decision tree model
// the Name field is the id of the instruction and have the same value as its corresponding model in the service response
// Criterion is the the name of the loss function that will be used to train this tree, supported values are 'RSS' and 'MSE'
// MinLeafSample indicates the min number of elements that have to be present in a leaf to stop tree splitting
type DecisiontreeRegIntruction struct {
	Name           string `json:"name"`
	Kind           string `json:"kind"`
	Criterion      string `json:"criterion"`
	MinLeafSamples int    `json:"minLeafSamples"`
}

// represents the training instructions for a clasification decision tree model
// the Name field is the id of the instruction and have the same value as its corresponding model in the service response
// Criterion is the the name of the loss function that will be used to train this tree, supported values are 'GINI' and 'ENTROPY'
type DecisiontreeClassIntruction struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Criterion string `json:"criterion"`
}

// parses the decision tree regression training instructions from a json file
// the json file must be a representation of an array of DecisiontreeRegIntruction
// for example :

// [
//   {
//     "name": "model3",
//     "kind": "regressor",
//     "criterion": "RSS",
//     "minLeafSamples": 20
//   },
//   {
//     "name": "model4",
//     "kind": "regressor",
//     "criterion": "MSE",
//     "minLeafSamples": 20
//   }
// ]
func ParseDTRegInstruction(f multipart.File) []DecisiontreeRegIntruction {
	filebytes, _ := ioutil.ReadAll(f)
	var instructions []DecisiontreeRegIntruction
	err := json.Unmarshal(filebytes, &instructions)
	if err != nil {
		log.Fatal(err)
	}

	return instructions
}

// parses the decision tree classifier training instructions from a json file
// the json file must be a representation of an array of DecisiontreeClassIntruction
// for example :

// [
//   {
//     "name": "model1",
//     "kind": "classifier",
//     "criterion": "GINI"
//   },
//   {
//     "name": "model2",
//     "kind": "classifier",
//     "criterion": "ENTROPY"
//   }
// ]
func ParseDTClassInstruction(f multipart.File) []DecisiontreeClassIntruction {
	filebytes, _ := ioutil.ReadAll(f)
	var instructions []DecisiontreeClassIntruction
	err := json.Unmarshal(filebytes, &instructions)
	if err != nil {
		log.Fatal(err)
	}

	return instructions
}
