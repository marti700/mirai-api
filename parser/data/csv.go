package data

import (
	"bufio"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/marti700/veritas/linearalgebra"
)

// Reads data from a csv file and returns the read data as a Matrix
// this functiona assumes the data in the csv are numbers in the float64 range
func ReadDataFromCSV(f multipart.File) linearalgebra.Matrix {

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
		matrixData += scanner.Text() + ","
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
	return linearalgebra.NewMatrix(matrix)
}
