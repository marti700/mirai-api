package reqhandler

import (
	"encoding/json"
	"fmt"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"

	"github.com/marti700/mirai/linearmodels"
	"github.com/marti700/mirai/options"
	"github.com/marti700/veritas/linearalgebra"
)

//holds the linear model response
// The ModelName field will have the same value that its corresponding instrucitions
// Model will be  the trained linear regression model
type LrResponse struct {
	ModelName string
	Model     linearmodels.LinearRegression
}

// Handles the requests made to the /regression endpoint
// this functions parses the training data and the training instructions (see the LrResponse struct)
// and responds with a json representation of the trained linear models
// for now data must be in csv format, the features and the target variable must be in defferent csv files
// instructions is a json file that indicates how the model must be trained
func HandleRegression(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	trainData := data.ReadDataFromCSV(dataFile)
	targetData := data.ReadDataFromCSV(targetFile)
	trainingInstructions := instruction.ParseInstruction(instructionsFile)
	resp, err := json.Marshal(trainModels(trainingInstructions, trainData, targetData))
	if err != nil {
		fmt.Println(err)
	}
	w.Write(resp)
}

func trainModels(trainIns []instruction.LinearRegInstructions, data, target linearalgebra.Matrix) []LrResponse {
	response := make([]LrResponse, len(trainIns))

	for i, ins := range trainIns {
		model := trainModel(ins, data, target)
		response[i] = LrResponse{
			ModelName: ins.Name,
			Model:     model,
		}

	}
	return response
}

// trains a linear regression model and returns the trained model
// if the OLS estimator is set to true then the gradiant descent estimator (GD) will be ignored
func trainModel(trainIns instruction.LinearRegInstructions, data, target linearalgebra.Matrix) linearmodels.LinearRegression {
	lr := linearmodels.LinearRegression{}

	if trainIns.Estimators.OLS {
		opts := options.LROptions{
			Estimator:      options.NewOLSOptions(),
			Regularization: trainIns.Regularization,
		}
		lr.Train(target, data, opts)
		return lr
	}

	if !isEmptyGD(trainIns.Estimators.GD) {
		opts := options.LROptions{
			Estimator:      trainIns.Estimators.GD,
			Regularization: trainIns.Regularization,
		}
		lr.Train(target, data, opts)
	}
	return lr
}

// Returns true if the GD options are not provided as an estimator to de linear regression model. Returns false otherwise
func isEmptyGD(gdOpts options.GDOptions) bool {
	if gdOpts.Iteations == 0 && gdOpts.LearningRate == 0 && gdOpts.MinStepSize == 0 {
		return true
	}
	return false
}
