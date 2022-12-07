package instruction

import (
	model "github.com/marti700/mirai/models"
	"github.com/marti700/mirai/models/linearmodels"
	"github.com/marti700/mirai/options"
)

func initalizeLinRegModel(trainIns []LinearRegInstructions) []model.Model {

	mod := make([]model.Model, len(trainIns))
	modelChanel := make(chan model.Model)

	for _, ins := range trainIns {
		go func(i LinearRegInstructions, ch chan model.Model) {
			lr := linearmodels.LinearRegression{}
			if i.Estimators.OLS {
				lr.Opts = options.LROptions{
					Estimator:      options.OLSOptions{},
					Regularization: i.Regularization,
				}
			} else {
				lr.Opts = options.LROptions{
					Estimator:      i.Estimators.GD,
					Regularization: i.Regularization,
				}
			}
			ch <- &lr
		}(ins, modelChanel)
	}

	for i := range mod {
		val := <-modelChanel
		mod[i] = val
	}

	return mod
}
