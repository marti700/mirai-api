package reqhandler

import (
	"mime/multipart"
	"net/http"
)

// get uploaded files from the request
func RequFiles(r *http.Request) (multipart.File, multipart.File, multipart.File, multipart.File, multipart.File) {
	instructionsFile, _, _ := r.FormFile("json")
	trainDataFile, _, _ := r.FormFile("trainData")
	trainTargetFile, _, _ := r.FormFile("trainTarget")
	testDataFile, _, _ := r.FormFile("testData")
	testTargetFile, _, _ := r.FormFile("testTarget")

	return instructionsFile, trainDataFile, trainTargetFile, testDataFile, testTargetFile
}

// close the provided files
func CloseFiles(f1, f2, f3, f4, f5 multipart.File) {
	f1.Close()
	f2.Close()
	f3.Close()
	f4.Close()
	f5.Close()
}

// returns the instructions in instructions.Instructions as a slice of LinearRegInstructions
// func coerceToLinRegIns(ins instruction.Instructions) []instruction.LinearRegInstructions {
// 	regIns := make([]instruction.LinearRegInstructions, len(ins.Instructions))

// 	for i, e := range ins.Instructions {
// 		regIns[i] = e.(instruction.LinearRegInstructions)
// 	}
// 	return regIns
// }
