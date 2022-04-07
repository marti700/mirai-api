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

func HandleRegression(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	trainData := data.ReadDataFromcsv(dataFile)
	targetData := data.ReadDataFromcsv(targetFile)
	trainingInstructions := instruction.ParseInstruction(instructionsFile, trainData, targetData)
	trainModel(trainingInstructions[1], trainData, targetData)
	fmt.Println(trainingInstructions)
}

func trainModel(trainIns instruction.LinearRegInstructions, data, target linearalgebra.Matrix) linearmodels.LinearRegression {
	lr := linearmodels.LinearRegression{}

	if trainIns.Estimators.OLS {
		opts := options.LROptions{
			Estimator: options.NewOLSOptions(),
			Regularization: trainIns.Regularization,
		}
		lr.Train(target, data, opts)
		return lr
	}

	if !emptyGD(trainIns.Estimators.GD) {
		opts := options.LROptions{
			Estimator: trainIns.Estimators.GD,
			Regularization: trainIns.Regularization,
		}
		lr.Train(target, data, opts)
	}
	return lr
}

func emptyGD(gdOpts options.GDOptions) bool {
	if gdOpts.Iteations == 0 && gdOpts.LearningRate == 0 && gdOpts.MinStepSize == 0 {
		return true
	}
	return false
}
