package reqhandler

import (
	"archive/zip"
	"io"
	"log"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/marti700/veritas/linearalgebra"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	instructions := instruction.Parse(instructionsFile)
	trainM(instructions, data.ReadDataFromCSV(dataFile), data.ReadDataFromCSV(targetFile))
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

func prepareFiles1(instructions []instruction.Instructions) {
	dirName := time.Now().String()
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// archive, err := os.Create(dirName + "/models.zip")

	for _, ins := range instructions {
		// create an empty zip file
		archive, err := os.Create(dirName + "/" + ins.InstructionType + ".zip")
		if err != nil {
			log.Fatal(err)
		}
		defer archive.Close()

		// createa a zip writer this will write files to the archive we created in line 52
		zipWriter := zip.NewWriter(archive)
		defer zipWriter.Close()
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

				// creates a file inside the zip archive, this is the file to be compresed
				zipFile, err := zipWriter.Create(ins.InstructionType + ".txt")
				if err != nil {
					log.Fatal(err)
				}

				// copies contents to the archive file
				io.Copy(zipFile, reportFile)
			}
		}
	}
}
