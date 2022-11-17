package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
)

// interface to be implemented by all the instructions
// the parse method parses the instructions from a json file
type Instruction interface {
	Parse(f multipart.File) []Instructions
}

// Represents a set of instructions
// the InstructionType field indicates what kind of struction this entity is
// the Name field is an id that will identify which model was trained with this instruction
// the Instructions field are the instructions used to train the model

type Instructions struct {
	InstructionType string        `json:modelType`
	Name            string        `json:name`
	Instructions    []Instruction `json:instructions`
}

// Creates and returns an empty Instructions Entity
func NewInstructions() Instructions {
	return Instructions{}
}

// Ej:

// [
//   {
//     "InstructionType": "linearregression",
//     "name": "first Instruction",
//     "instructions": {
//       "estimators": {
//         "GD": {
//           "Iteations": 1000,
//           "LearningRate": 0.01,
//           "MinStepSize": 0.00003
//         },
//         "OLS": false
//       },
//       "regularization": {
//         "type": "l1",
//         "lambda": 20.0
//       }
//     }
//   },
//   {
//     "InstructionType": "DecisionTreeClassifier",
//     "name": "second Instruction DTC",
//     "instructions": {
//       "name": "Instruction1",
//       "kind": "classifier",
//       "criterion": "GINI"
//     }
//   },
//   {
//     "InstructionType": "DecisionTreeRegressor",
//     "name": "third Instruction",
//     "instructions": {
//       "name": "Instruction3",
//       "kind": "regressor",
//       "criterion": "RSS",
//       "minLeafSamples": 20
//     }
//   }
// ]

func (i Instructions) Parse(f multipart.File) []Instructions {
	filebytes, _ := ioutil.ReadAll(f)
	var ins []Instructions
	// err := json.Unmarshal(filebytes, &ins)
	var objmap []interface{} // loop through this
	// var objmap map[string]interface{} // declare this in neach iteration of the loop
	// objmap["instructionType"] and objmap["instrctions"]

	err := json.Unmarshal(filebytes, &objmap)
	// err1 := json.Unmarshal(r, &li)

	if err != nil {
		log.Fatal(err)
	}

	return ins
}
