package main

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"mirai-api/instruction"
	"net/http"
	"strconv"
	"strings"

	"github.com/marti700/veritas/linearalgebra"
)

// Reads data from a csv file and returns the read data as a Matrix
// this functiona assumes the data in the csv are numbers in the float64 range
func ReadDataFromcsv(f multipart.File) linearalgebra.Matrix {

	scanner := bufio.NewScanner(f)
	var matrixData string

	// read file first line to get the matrix column number
	// this are the heading numbers of the csv files
	// this line can be discarted since it does not hold useful data
	scanner.Scan()
	fstLine := scanner.Text()
	cols_num := len(strings.Split(fstLine, ","))

	// loop through the rest of the file
	fileLines := 0
	for scanner.Scan() {
		// extra coma so that the last number of this line don't get mixed with the first number of the next when slitting later
		matrixData += scanner.Text()+","
		fileLines++
	}

	dataSplit := strings.Split(matrixData, ",")

	matrix := make([][]float64, fileLines)
	row := 0
	col := 0
	nextBreak := cols_num

	data := make([]float64, cols_num)
	for i, e := range dataSplit {
		//before processing the matrix next row
		if i == nextBreak {
			matrix[row] = data
			data = make([]float64, cols_num)
			nextBreak += cols_num
			row++
			col = 0
		}
		data[col], _ = strconv.ParseFloat(e, 64)
		col++
	}
	fmt.Println(linearalgebra.NewMatrix(matrix))
	return linearalgebra.NewMatrix(matrix)
}

func HandleUpload(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(200)
	instructions, handler, _ := r.FormFile("json")
	data, handler, _ := r.FormFile("train")
	target, handler, _ := r.FormFile("target")
	defer instructions.Close()
	defer data.Close()
	defer target.Close()
	trainData := ReadDataFromcsv(data)
	targetData := ReadDataFromcsv(target)
	instruction.ParseInstruction(instructions, trainData, targetData)
	fmt.Println(handler.Filename)
	fmt.Println(handler.Size)
	fmt.Println("ja")

}

func main (){
	http.HandleFunc("/upload", HandleUpload)
	http.ListenAndServe(":9090", nil)
}