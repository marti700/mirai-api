package reqhandler

import (
	"net/http"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
)
func HandleRegression(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(200)
	instructionsFile, _, _ := r.FormFile("json")
	dataFile, _, _ := r.FormFile("train")
	targetFile, _, _ := r.FormFile("target")
	defer instructionsFile.Close()
	defer dataFile.Close()
	defer targetFile.Close()
	trainData := data.ReadDataFromcsv(dataFile)
	targetData := data.ReadDataFromcsv(targetFile)
	instruction.ParseInstruction(instructionsFile, trainData, targetData)
}
