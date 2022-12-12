package reqhandler

import (
	"fmt"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"
	"sync"

	model "github.com/marti700/mirai/models"
	"github.com/marti700/veritas/linearalgebra"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	instructions := instruction.Parse(instructionsFile)
	trainM(instructions, data.ReadDataFromCSV(dataFile), data.ReadDataFromCSV(targetFile))
	fmt.Println(instructions)
}

func trainM(trainInstructions []instruction.Instructions, data, target linearalgebra.Matrix) {
	var wg sync.WaitGroup
	for _, instruction := range trainInstructions {
		for _, models := range instruction.Models {
			for key, value := range models {
				wg.Add(1)
				go func(key string, mod model.Model) {
					mod.Train(data, target)
					wg.Done()
				}(key, value)
			}
		}
	}
	wg.Wait()
}
