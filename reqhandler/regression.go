package reqhandler

import (
	"fmt"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"

	"github.com/marti700/mirai/linearmodels"
	"github.com/marti700/mirai/options"
	"github.com/marti700/veritas/linearalgebra"
)

// Handles the requests made to the /regression endpoint
// this functions parses the training data and the training instructions
// and responds with a json representation of the trained linear model (TODO)
// for now data must be in csv format, the features and the target variable mustbe in defferent csv files
// instructions is a json file that indicates how the data must be trained
func HandleRegression(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	trainData := data.ReadDataFromCSV(dataFile)
	targetData := data.ReadDataFromCSV(targetFile)
	trainingInstructions := instruction.ParseInstruction(instructionsFile, trainData, targetData)
	trainModel(trainingInstructions[1], trainData, targetData)
	fmt.Println(trainingInstructions)
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

	if !emptyGD(trainIns.Estimators.GD) {
		opts := options.LROptions{
			Estimator:      trainIns.Estimators.GD,
			Regularization: trainIns.Regularization,
		}
		lr.Train(target, data, opts)
	}
	return lr
}

// Returns true if the GD options are not provided as an estimator to de linear regression model. Returns false otherwise
func emptyGD(gdOpts options.GDOptions) bool {
	if gdOpts.Iteations == 0 && gdOpts.LearningRate == 0 && gdOpts.MinStepSize == 0 {
		return true
	}
	return false
}
