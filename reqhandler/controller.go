package reqhandler

import (
	"archive/zip"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"mirai-api/report"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/marti700/veritas/linearalgebra"
)

type response struct {
	Status       string
	ErrorMessage string
}

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response{}
	instructionsFile, trainDataFile, trainTargetFile, testDataFile, testTargetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, trainDataFile, trainTargetFile, testDataFile, testTargetFile)

	instructions := instruction.Parse(instructionsFile)
	trainM(instructions, data.ReadDataFromCSV(trainDataFile), data.ReadDataFromCSV(trainTargetFile),
		data.ReadDataFromCSV(testDataFile), data.ReadDataFromCSV(testTargetFile))
	reportsDirectory, err := prepareReports(instructions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.ErrorMessage = err.Error()
		resp.Status = "Fail"
		json.NewEncoder(w).Encode(resp)
		return
	}
	reportFiles := reportsDirectory + "reports.zip"

	err = report.SendReportByEmail(r.URL.Query()["email"][0], reportFiles)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.ErrorMessage = err.Error()
		resp.Status = "Fail"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// deletes the reports after sending them by email
	err = os.RemoveAll(reportsDirectory)
	if err != nil {
		log.Print(err)
	}

	resp.Status = "OK"
	resp.ErrorMessage = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// trains models based on the provided instructions
func trainM(trainInstructions []instruction.Instructions, data, target, testData, testTarget linearalgebra.Matrix) {
	var wg sync.WaitGroup
	for _, ins := range trainInstructions {
		for _, models := range ins.Models {
			for key, value := range models {
				wg.Add(1)
				go func(key string, mod instruction.MiraiModel) {
					mod.Mod.Train(data, target)
					predictions := mod.Mod.Predict(testData)
					mod.Report.CreateReport(testTarget, predictions, mod.Mod)
					mod.Report.ToString()
					wg.Done()
				}(key, value)
			}
		}
	}
	wg.Wait()
}

// creates the reports that will be send by email
func prepareReports(instructions []instruction.Instructions) (string, error) {
	dirName := strconv.FormatInt(time.Now().Unix(), 10)
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return "", err
	}

	for _, ins := range instructions {
		for _, mods := range ins.Models {
			for key, value := range mods {
				reportString := key + " \n" + value.Report.ToString() + "\n"
				reportsPath := dirName + "/" + ins.InstructionType + ".txt"
				// creates the report file, if the file does not exists create it otherwise just append data
				reportFile, err := os.OpenFile(reportsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return "", err
				}

				reportFile.WriteString(reportString)
				defer reportFile.Close()

			}
		}
	}
	err = zipReports(dirName)
	if err != nil {
		return "", err
	}

	return dirName + "/", nil
}

// puts the specified directory into a zip file
func zipReports(pathToReportFolder string) error {
	// create an empty zip file
	archive, err := os.Create(pathToReportFolder + "/reports" + ".zip")
	if err != nil {
		return err
	}
	defer archive.Close()

	// createa a zip writer this will write files to the archive we created in line 52
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	walkerFunc := func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// avoids the zip file from being processing and prevents the root folder name to be taken into consideration
		if !strings.HasSuffix(path, "zip") {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			// creates a file inside the zip archive, this is the file to be compresed
			if path != pathToReportFolder {
				zipFile, err := zipWriter.Create(path)
				if err != nil {
					return err
				}
				// copies contents to the archive file
				io.Copy(zipFile, f)
			}
		}
		return nil
	}

	err = filepath.Walk(pathToReportFolder, walkerFunc)
	if err != nil {
		return err
	}

	return nil
}
