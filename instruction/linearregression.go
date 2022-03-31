package instruction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"

	// "github.com/marti700/mirai/linearmodels"
	"github.com/marti700/mirai/options"
	"github.com/marti700/veritas/linearalgebra"
)

type estimators struct {
	GD  []options.GDOptions `json:"GD"`
	OLS bool                `json:"OLS"`
}

type LinearRegInstructions struct {
	Estimators     estimators         `json:"estimators"`
	Regularization []options.RegOptions `json:"regularization"`
}

func ParseInstruction(f multipart.File, data, target linearalgebra.Matrix) {
	filebytes, _ := ioutil.ReadAll(f)
	linRegInstructions := LinearRegInstructions{}
	err := json.Unmarshal(filebytes, &linRegInstructions)
	if err != nil {
		log.Fatal(err)
	}

	// lr := linearmodels.LinearRegression{}

	// opts := options.LROptions{
	// 	Estimator: linRegInstructions.Estimator.Ols,
	// }

	// lr.Train(target, data, opts)
}