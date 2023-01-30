package reqhandler

import (
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"
	"sync"

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
					wg.Done()
				}(key, value)
			}
		}
	}
	wg.Wait()
}

func generateReports(trainInstructions []instruction.Instructions, data, target linearalgebra.Matrix) {
	// var wg sync.WaitGroup
	// for _, ins := range trainInstructions {
	// 	for _, models := range ins.Models {
	// 		for key, value := range models {
	// 			wg.Add(1)
	// 			go func(key string, mod instruction.MiraiModel) {
	// 				rep := mod.Report
	// 				mod.Report = rep.CreateReport(data, target, mod.Mod)
	// 				wg.Done()
	// 			}(key, value)
	// 		}
	// 	}
	// }
	// wg.Wait()
}
