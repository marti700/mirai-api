package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"

	model "github.com/marti700/mirai/models"
)

// interface to be implemented by all the instructions
// the parse method parses the instructions from a json file
type Instruction interface {
	Parse(f multipart.File) Instructions
}

// Represents a set of instructions
// the InstructionType field indicates what kind of struction this entity is
// the Name field is an id that will identify which model was trained with this instruction
// the Instructions field are the instructions used to train the model

type Instructions struct {
	InstructionType string        `json:modelType`
	Name            string        `json:name`
	Models          []model.Model `json:instructions`
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
	inss := make([]Instructions, 0, 200)
	// mod := make([]model.Model, 200)

	var objmap []interface{}
	err := json.Unmarshal(filebytes, &objmap)
	// j := objmap[0]
	// r := j.(map[string]interface{})

	for _, ins := range objmap {
		instruction := ins.(map[string]interface{})
		insType := instruction["InstructionType"].(string)
		if insType == "linearregression" {
			//create model initializers package to initialize model depending on their types ej: linRegInitializer, TreeInitializer, etc
			linRegIns := getLinearRegIns(instruction["instructions"].([]interface{}))
			// copy(mod, initalizeLinRegModel(linRegIns))

			regModels := Instructions{
				InstructionType: insType,
				Name:            instruction["name"].(string),
				Models:          initalizeLinRegModel(linRegIns),
			}
			inss = append(inss, regModels)
		}
	}
	// err1 := json.Unmarshal(r, &li)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(r)

	return inss
}

func getLinearRegIns(ins []interface{}) []LinearRegInstructions {
	var linRegIns []LinearRegInstructions
	// x := h.([]interface{})
	b, _ := json.Marshal(ins)
	// b := x[0].([]byte)
	error := json.Unmarshal(b, &linRegIns)

	if error != nil {
		log.Fatal(error)
	}

	return linRegIns
}
