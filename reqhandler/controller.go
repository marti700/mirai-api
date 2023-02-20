package reqhandler

import (
	"archive/zip"
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

func Handle(w http.ResponseWriter, r *http.Request) {
	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	instructions := instruction.Parse(instructionsFile)
	trainM(instructions, data.ReadDataFromCSV(dataFile), data.ReadDataFromCSV(targetFile))
	reportsDirectory := prepareReports(instructions)
	reportFiles := reportsDirectory + "reports.zip"

	report.SendReportByEmail(r.URL.Query()["email"][0], reportFiles)

	// deletes the reports after sending them by email
	err := os.RemoveAll(reportsDirectory)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func trainM(trainInstructions []instruction.Instructions, data, target linearalgebra.Matrix) {
	var wg sync.WaitGroup
	for _, ins := range trainInstructions {
		for _, models := range ins.Models {
			for key, value := range models {
				wg.Add(1)
				go func(key string, mod instruction.MiraiModel) {
					mod.Mod.Train(data, target)
					mod.Report.CreateReport(data, target, mod.Mod)
					mod.Report.ToString()
					wg.Done()
				}(key, value)
			}
		}
	}
	wg.Wait()
}

func prepareReports(instructions []instruction.Instructions) string {
	dirName := strconv.FormatInt(time.Now().Unix(), 10)
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, ins := range instructions {
		for _, mods := range ins.Models {
			for key, value := range mods {
				reportString := key + " \n" + value.Report.ToString() + "\n"
				reportsPath := dirName + "/" + ins.InstructionType + ".txt"
				// creates the report file, if the file does not exists create it otherwise just append data
				reportFile, err := os.OpenFile(reportsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}

				reportFile.WriteString(reportString)
				defer reportFile.Close()

			}
		}
	}
	zipReports(dirName)

	return dirName + "/"
}

func zipReports(pathToReportFolder string) {
	// create an empty zip file
	archive, err := os.Create(pathToReportFolder + "/reports" + ".zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	// createa a zip writer this will write files to the archive we created in line 52
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	walkerFunc := func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// avoids the zip file processing and avoids the root folder name to be taken into consideration
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
		log.Fatal(err)
	}
}
