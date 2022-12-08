package instruction

import (
	"github.com/marti700/mirai/metrics"
	model "github.com/marti700/mirai/models"
	"github.com/marti700/mirai/models/linearmodels"
	"github.com/marti700/mirai/models/treemodels"
	"github.com/marti700/mirai/options"
)

// Recieves a LinearRegressionInstructions slice to initialize linear regression models
func initalizeLinRegModel(trainIns []LinearRegInstructions) []map[string]model.Model {
	mod := make([]map[string]model.Model, len(trainIns))
	// modelChanel := make(chan model.Model)

	for i, ins := range trainIns {
		lr := linearmodels.LinearRegression{}
		if ins.Estimators.OLS {
			lr.Opts = options.LROptions{
				Estimator:      options.OLSOptions{},
				Regularization: ins.Regularization,
			}
		} else {
			lr.Opts = options.LROptions{
				Estimator:      ins.Estimators.GD,
				Regularization: ins.Regularization,
			}
		}
		m := make(map[string]model.Model)
		m[ins.Name] = &lr
		mod[i] = m
	}

	return mod
}

func initializeDecisionTreeClassifierModel(trainIns []DecisiontreeClassIntruction) []map[string]model.Model {
	mod := make([]map[string]model.Model, len(trainIns))

	//Initialize models
	for i, ins := range trainIns {
		c_opt := options.NewDTreeClassifierOption(ins.Criterion)
		treeModel := treemodels.NewDecicionTreeeClassifier(c_opt)

		m := make(map[string]model.Model)
		m[ins.Name] = treeModel
		mod[i] = m
	}

	return mod
}

func initializeDecisionTreeRegressorModel(trainIns []DecisiontreeRegIntruction) []map[string]model.Model {
	mod := make([]map[string]model.Model, len(trainIns))

	//Initialize models
	for i, ins := range trainIns {
		var reg_opt options.DTreeRegreessorOptions
		if ins.Criterion == "RSS" {
			reg_opt = options.NewDTRegressorOptions(ins.MinLeafSamples, metrics.RSS)
		} else if ins.Criterion == "MSE" {
			reg_opt = options.NewDTRegressorOptions(ins.MinLeafSamples, metrics.MeanSquareError)
		}

		treeModel := treemodels.NewDecisionTreeRegressor(reg_opt)

		m := make(map[string]model.Model)
		m[ins.Name] = treeModel
		mod[i] = m
	}

	return mod
}
