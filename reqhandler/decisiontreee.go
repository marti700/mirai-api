package reqhandler

import (
	"archive/zip"
	"time"
	// "encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	// "log"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"

	// "time"

	"github.com/marti700/mirai/metrics"
	model "github.com/marti700/mirai/models"
	"github.com/marti700/mirai/models/treemodels"
	"github.com/marti700/mirai/options"
	"github.com/marti700/veritas/linearalgebra"
)

type DTResponse struct {
	ModelName string
	Model     model.Model
}

func newDTResponse(id string, mod model.Model) DTResponse {
	return DTResponse{
		ModelName: id,
		Model:     mod,
	}
}

// Since a tree can be hard to read as a text instead of a json response
// a .zip file with the models is returned as application/octet-stream
// the zip file contains the models as .dot files
// (https://en.wikipedia.org/wiki/DOT_(graph_description_language))
func HandleDecisionTree(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	trainData := data.ReadDataFromCSV(dataFile)
	targetData := data.ReadDataFromCSV(targetFile)
	trainingInstructions := instruction.ParseInstruction1(instructionsFile)
	filePath:= prepareFiles(trainDecisionTree(trainingInstructions, trainData, targetData))

	f, err := ioutil.ReadFile(filePath+"/models.zip")
	if err != nil {
		log.Fatal(err)
	}

	// resp, err := json.Marshal(trainDecisionTree(trainingInstructions, trainData, targetData))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(f)

	err = os.RemoveAll(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func trainDecisionTree(trainingInstructions []instruction.DecisiontreeIntruction,
	data, target linearalgebra.Matrix) []DTResponse {

	resp := classifyModels(trainingInstructions)
	trainDTs(data, target, resp)

	return resp
}

func classifyModels(ins []instruction.DecisiontreeIntruction) []DTResponse {
	responses := make([]DTResponse, 0, len(ins))

	for _, m := range ins {
		if m.Kind == "regressor" {
			var reg_opt options.DTreeRegreessorOptions
			if m.Criterion == "RSS" {
				reg_opt = options.NewDTRegressorOptions(m.MinLeafSamples, metrics.RSS)
			} else if m.Criterion == "MSS" {
				reg_opt = options.NewDTRegressorOptions(m.MinLeafSamples, metrics.MeanSquareError)
			}
			model := treemodels.NewDecisionTreeRegressor(reg_opt)
			resp := newDTResponse(m.Name, model)
			responses = append(responses, resp)
		} else {
			c_opt := options.NewDTreeClassifierOption(m.Criterion)
			model := treemodels.NewDecicionTreeeClassifier(c_opt)
			resp := newDTResponse(m.Name, model)
			responses = append(responses, resp)
		}
	}

	return responses
}

// trains the models present in a DTResponse
// and retunrs a new DTResponse array with trained models
func trainDTs(data, target linearalgebra.Matrix,
	models []DTResponse) []DTResponse {

		resps := make([]DTResponse, len(models))

	for i, m := range models {
		m.Model.Train(data, target)
		resps[i] = m
	}

	return resps
}

// given a DTResponse array creates a zip file containing
// the models as a .dot files
// returns the location where the zip file is
func prepareFiles(models []DTResponse) string {
dirName := time.Now().String()
	os.Mkdir(dirName, os.ModePerm)
	archive, err := os.Create(dirName + "/models.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

for _, m := range models {
		switch t := m.Model.(type) {
		case *treemodels.DecisionTreeClassifier:
			zwc, _ := zipWriter.Create(m.ModelName + ".dot")
			t.Model.Plot()
			f, _ := os.Open("tree.dot")
			io.Copy(zwc, f)
			f.Close()
		case *treemodels.DecisionTreeRegressor:
			zwr, _ := zipWriter.Create(m.ModelName + ".dot")
			t.Model.Plot()
			f, _ := os.Open("tree.dot")
			io.Copy(zwr, f)
			f.Close()
		}
	}
	return dirName
}
