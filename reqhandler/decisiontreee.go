package reqhandler

import (
	"archive/zip"
	"sync"
	"time"

	"io"
	"io/ioutil"
	"log"
	"os"

	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"

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
func HandleDecisionTreeRegressor(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile, testData, testTarget := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile, testData, testTarget)

	trainData := data.ReadDataFromCSV(dataFile)
	targetData := data.ReadDataFromCSV(targetFile)

	trainingInstructions := instruction.NewDecisionRegresor().Parse(instructionsFile)
	filePath := prepareFiles(trainDecisionTreeRegressor(trainingInstructions, trainData, targetData))

	f, err := ioutil.ReadFile(filePath + "/models.zip")
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(f)

	err = os.RemoveAll(filePath) // remove temporary files
	if err != nil {
		log.Fatal(err)
	}
}

// Since a tree can be hard to read as a text instead of a json response
// a .zip file with the models is returned as application/octet-stream
// the zip file contains the models as .dot files
// (https://en.wikipedia.org/wiki/DOT_(graph_description_language))
func HandleDecisionTreeClassifier(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200)

	// instructionsFile, dataFile, targetFile := RequFiles(r)
	instructionsFile, dataFile, targetFile, testData, testTarget := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile, testData, testTarget)

	trainData := data.ReadDataFromCSV(dataFile)
	targetData := data.ReadDataFromCSV(targetFile)
	trainingInstructions := instruction.NewDecisionClassifier().Parse(instructionsFile)
	filePath := prepareFiles(trainDecisionTreeClassifier(trainingInstructions, trainData, targetData))

	f, err := ioutil.ReadFile(filePath + "/models.zip")
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(f)

	// removes temporary files
	err = os.RemoveAll(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

// trains the regression models present in a DTResponse
// and retunrs a new DTResponse array with the trained models
func trainDecisionTreeRegressor(trainingInstructions []instruction.DecisiontreeRegIntruction,
	data, target linearalgebra.Matrix) []DTResponse {

	responses := make([]DTResponse, 0, len(trainingInstructions))

	//Initialize models
	for _, m := range trainingInstructions {
		var reg_opt options.DTreeRegreessorOptions
		if m.Criterion == "RSS" {
			reg_opt = options.NewDTRegressorOptions(m.MinLeafSamples, metrics.RSS)
		} else if m.Criterion == "MSE" {
			reg_opt = options.NewDTRegressorOptions(m.MinLeafSamples, metrics.MeanSquareError)
		}

		model := treemodels.NewDecisionTreeRegressor(reg_opt)
		resp := newDTResponse(m.Name, model)
		responses = append(responses, resp)
	}

	// train models concurrently
	var wg sync.WaitGroup

	for _, m := range responses {
		wg.Add(1)
		go func(mod DTResponse) {
			defer wg.Done()
			mod.Model.Train(data, target)
		}(m)
	}
	wg.Wait()

	return responses
}

// trains the classification models present in a DTResponse
// and retunrs a new DTResponse array with trained models
func trainDecisionTreeClassifier(trainingInstructions []instruction.DecisiontreeClassIntruction,
	data, target linearalgebra.Matrix) []DTResponse {

	responses := make([]DTResponse, 0, len(trainingInstructions))

	//Initialize models
	for _, m := range trainingInstructions {

		c_opt := options.NewDTreeClassifierOption(m.Criterion)
		model := treemodels.NewDecicionTreeeClassifier(c_opt)
		resp := newDTResponse(m.Name, model)
		responses = append(responses, resp)
	}

	// train models concurrently
	var wg sync.WaitGroup

	for _, m := range responses {
		wg.Add(1)
		go func(mod DTResponse) {
			defer wg.Done()
			mod.Model.Train(data, target)
		}(m)
	}
	wg.Wait()

	return responses
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
