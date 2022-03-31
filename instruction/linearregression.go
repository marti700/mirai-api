package instruction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/marti700/mirai/linearmodels"
	"github.com/marti700/mirai/options"
	"github.com/marti700/veritas/linearalgebra"
)

type linearRegInstructions struct {
	Estimator      options.GDOptions  `json:estimator`
	Regularization options.RegOptions `json:regularization`
}

func ParseInstruction(f multipart.File, data,target linearalgebra.Matrix) {
	filebytes, _ := ioutil.ReadAll(f)
	linRegInstructions := linearRegInstructions{}
	err := json.Unmarshal(filebytes, &linRegInstructions)
	if err != nil{
		log.Fatal(err)
	}

	lr := linearmodels.LinearRegression{}

	opts := options.LROptions{
		Estimator: linRegInstructions.Estimator,
	}

	lr.Train(target, data, opts)
}
