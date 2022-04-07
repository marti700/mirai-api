package reqhandler

import (
	"fmt"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"net/http"
)
func HandleRegression(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(200)

	instructionsFile, dataFile, targetFile := RequFiles(r)
	defer CloseFiles(instructionsFile, dataFile, targetFile)

	trainData := data.ReadDataFromcsv(dataFile)
	targetData := data.ReadDataFromcsv(targetFile)
	trainingInstructions := instruction.ParseInstruction(instructionsFile, trainData, targetData)
	fmt.Println(trainingInstructions)
}
