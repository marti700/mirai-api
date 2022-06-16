package reqhandler

import (
	"archive/zip"
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
	// classifier *treemodels.DecisionTreeClassifier
	// regressor   *treemodels.DecisionTreeRegressor
}

func newDTResponse(id string, mod model.Model) DTResponse {
	return DTResponse{
		ModelName: id,
		Model:     mod,
	}
}

func HandleDecisionTree(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	trainData := data.ReadDataFromCSV(dataFile)
	targetData := data.ReadDataFromCSV(targetFile)
	trainingInstructions := instruction.ParseInstruction1(instructionsFile)
	trainDecisionTree(trainingInstructions, trainData, targetData)

	f, err := ioutil.ReadFile("models.zip")
	if err != nil {
		log.Fatal("Error preparing models download")
	}

	// resp, err := json.Marshal(trainDecisionTree(trainingInstructions, trainData, targetData))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(f)
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

func trainDTs(data, target linearalgebra.Matrix,
	models []DTResponse) {

	archive, err := os.Create("models.zip")
	if err != nil {
		log.Fatal("Error creating zip file")
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for _, m := range models {
		m.Model.Train(data, target)
		filePath := "./" + m.ModelName+".dot"
		dst, _ := os.Create(filePath)
		switch t := m.Model.(type) {
		case *treemodels.DecisionTreeClassifier:
			t.Model.Plot()
			f, _ := os.Open("tree.dot")
			zipWriter.Create(m.ModelName)
			io.Copy(dst, f)
		case *treemodels.DecisionTreeRegressor:
			t.Model.Plot()
			f, _ := os.Open("tree.dot")
			zipWriter.Create(m.ModelName)
			io.Copy(dst, f)
		}
	}
}
