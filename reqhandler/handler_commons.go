package reqhandler

import (
	"mime/multipart"
	"net/http"
)

// get uploaded files from the request
func RequFiles(r *http.Request) (multipart.File, multipart.File, multipart.File) {
	instructionsFile, _, _ := r.FormFile("json")
	dataFile, _, _ := r.FormFile("train")
	targetFile, _, _ := r.FormFile("target")

	return instructionsFile, dataFile, targetFile
}

// close the provided files
func CloseFiles(f1, f2, f3 multipart.File) {
	f1.Close()
	f2.Close()
	f3.Close()
}
