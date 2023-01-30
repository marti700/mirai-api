package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"mirai-api/report"

	model "github.com/marti700/mirai/models"
)

type MiraiModel struct {
	Mod    model.Model
	Report report.Reporter
}

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
	InstructionType string                  `json:"modelType"`
	Name            string                  `json:"name"`
	Models          []map[string]MiraiModel `json:"instructions"`
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
//     "InstructionType": "decisiontreeclassifier",
//     "name": "second Instruction DTC",
//     "instructions": {
//       "name": "Instruction1",
//       "kind": "classifier",
//       "criterion": "GINI"
//     }
//   },
//   {
//     "InstructionType": "decisiontreeregressor",
//     "name": "third Instruction",
//     "instructions": {
//       "name": "Instruction3",
//       "kind": "regressor",
//       "criterion": "RSS",
//       "minLeafSamples": 20
//     }
//   }
// ]

func Parse(f multipart.File) []Instructions {
	filebytes, _ := ioutil.ReadAll(f)
	inss := make([]Instructions, 0, 200)

	var objmap []interface{}
	err := json.Unmarshal(filebytes, &objmap)

	for _, ins := range objmap {
		instruction := ins.(map[string]interface{})
		insType := instruction["InstructionType"].(string)
		if insType == "linearregression" {
			//create model initializers package to initialize model depending on their types ej: linRegInitializer, TreeInitializer, etc
			linRegIns := getLinearRegIns(instruction["instructions"].([]interface{}))

			regModels := Instructions{
				InstructionType: insType,
				Name:            instruction["name"].(string),
				Models:          initalizeLinRegModel(linRegIns),
			}
			inss = append(inss, regModels)
		}

		if insType == "decisiontreeclassifier" {
			dtc := getDTCIns(instruction["instructions"].([]interface{}))

			dtcModel := Instructions{
				InstructionType: insType,
				Name:            instruction["name"].(string),
				Models:          initializeDecisionTreeClassifierModel(dtc),
			}

			inss = append(inss, dtcModel)
		}
		if insType == "decisiontreeregressor" {
			dtr := getDTRegIns(instruction["instructions"].([]interface{}))

			dtrModel := Instructions{
				InstructionType: insType,
				Name:            instruction["name"].(string),
				Models:          initializeDecisionTreeRegressorModel(dtr),
			}

			inss = append(inss, dtrModel)

		}
	}

	if err != nil {
		log.Fatal(err)
	}

	return inss
}

func getLinearRegIns(ins []interface{}) []LinearRegInstructions {
	var linRegIns []LinearRegInstructions
	// marshals the ins parameter to get an array of bytes
	// it will be used later to convert it to an LinearRegInstructions
	b, _ := json.Marshal(ins)
	error := json.Unmarshal(b, &linRegIns)

	if error != nil {
		log.Fatal(error)
	}

	return linRegIns
}

func getDTCIns(ins []interface{}) []DecisiontreeClassIntruction {
	var dtcIns []DecisiontreeClassIntruction
	b, _ := json.Marshal(ins)
	error := json.Unmarshal(b, &dtcIns)

	if error != nil {
		log.Fatal(error)
	}

	return dtcIns
}

func getDTRegIns(ins []interface{}) []DecisiontreeRegIntruction {
	var dtrIns []DecisiontreeRegIntruction
	b, _ := json.Marshal(ins)
	error := json.Unmarshal(b, &dtrIns)

	if error != nil {
		log.Fatal(error)
	}

	return dtrIns
}
