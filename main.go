package main

import (
	"fmt"
	"net/http"

	"mirai-api/parser/data"
	"mirai-api/parser/instruction"

)



func HandleUpload(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(200)
	instructionsFile, _, _ := r.FormFile("json")
	dataFile, _, _ := r.FormFile("train")
	targetFile, handler, _ := r.FormFile("target")
	defer instructionsFile.Close()
	defer dataFile.Close()
	defer targetFile.Close()
	trainData := data.ReadDataFromcsv(dataFile)
	targetData := data.ReadDataFromcsv(targetFile)
	instruction.ParseInstruction(instructionsFile, trainData, targetData)
	fmt.Println(handler.Filename)
	fmt.Println(handler.Size)
	fmt.Println("ja")

}

func main (){
	http.HandleFunc("/upload", HandleUpload)
	http.ListenAndServe(":9090", nil)
}