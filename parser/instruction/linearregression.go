package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/marti700/mirai/options"
	// "github.com/marti700/veritas/linearalgebra"
)

// consists of the suported model estimators of the mirai linear regression model
type Estimators struct {
	GD  options.GDOptions `json:"GD"`
	OLS bool              `json:"OLS"`
}


// represents a linear regression instructions
// the Name field is the id of the instruction and have the same value as its corresponding model in the service response
// Estimators are the linear model estimators supported by mirai current options are GD which will train the model using Gradiant Descent
// and OLS which estimates the model hyperparameters from the linear regression close form solution
// the regularization parameter specifies if the model should be trained using regularization. Supported values are l1 (lasso) and l2 (ridge)
type LinearRegInstructions struct {
	Name           string             `json:"name"`
	Estimators     Estimators         `json:"estimators"`
	Regularization options.RegOptions `json:"regularization"`
}

// parses the linear regression training instructions from a json file
// the json file must be a representation of an array of LinearRegInstructions
// for example :

// [
//   {
//     "name": "first model",
//     "estimators": {
//       "GD": {
//         "Iteations": 1000,
//         "LearningRate": 0.001,
//         "MinStepSize": 0.00003
//       },
//       "OLS": true
//     },
//     "regularization": {
//       "type": "l1",
//       "lambda": 0.01
//     }
//   },
//   {
//     "name": "second model",
//     "estimators": {
//       "GD": {
//         "Iteations": 100,
//         "LearningRate": 0.01,
//         "MinStepSize": 0.002
//       },
//       "OLS": false
//     },
//     "regularization": {
//       "type": "l2",
//       "lambda": 0.01
//     }
//   },
//   {
//     "name": "third model",
//     "estimators": {
//       "OLS": true
//     }
//   }
// ]
func ParseInstruction(f multipart.File) []LinearRegInstructions {
	filebytes, _ := ioutil.ReadAll(f)
	var linRegInstructions []LinearRegInstructions
	err := json.Unmarshal(filebytes, &linRegInstructions)
	if err != nil {
		log.Fatal(err)
	}

	return linRegInstructions
}
