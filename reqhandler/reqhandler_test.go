package reqhandler

import (
	"fmt"
	"mirai-api/parser/data"
	"mirai-api/parser/instruction"
	"os"
	"testing"
)

func BenchmarkTrainModels(b *testing.B) {
	// parser.ReadDataFromCSV()
	trainDataFile, _ := os.Open("./benchmarkdata/x_train.csv")
	targetDataFile, _ := os.Open("./benchmarkdata/y_train.csv")
	instructionFile, _ := os.Open("./benchmarkdata/linReg.json")
	defer trainDataFile.Close()
	defer targetDataFile.Close()
	defer instructionFile.Close()

	train := data.ReadDataFromCSV(trainDataFile)
	target := data.ReadDataFromCSV(targetDataFile)
	instructions := instruction.NewLinearRegInstructions().Parse(instructionFile)

	for i := 0; i < b.N; i++ {
		trainModels(instructions, train, target)
	}

}

func BenchmarkTrainDecisionTreeRegressor(b *testing.B) {
	// parser.ReadDataFromCSV()
	trainDataFile, _ := os.Open("./benchmarkdata/x_train.csv")
	targetDataFile, _ := os.Open("./benchmarkdata/y_train.csv")
	instructionFile, _ := os.Open("./benchmarkdata/decisionTreeRegressor.json")
	defer trainDataFile.Close()
	defer targetDataFile.Close()
	defer instructionFile.Close()

	train := data.ReadDataFromCSV(trainDataFile)
	target := data.ReadDataFromCSV(targetDataFile)
	instructions := instruction.NewDecisionRegresor().Parse(instructionFile)

	for i := 0; i < b.N; i++ {
		trainDecisionTreeRegressor(instructions, train, target)
	}

}

func BenchmarkTrainDecisionTreeClassier(b *testing.B) {
	// parser.ReadDataFromCSV()
	trainDataFile, _ := os.Open("./benchmarkdata/x_train.csv")
	targetDataFile, _ := os.Open("./benchmarkdata/y_train.csv")
	instructionFile, _ := os.Open("./benchmarkdata/decisionTreeClassifier.json")
	defer trainDataFile.Close()
	defer targetDataFile.Close()
	defer instructionFile.Close()

	train := data.ReadDataFromCSV(trainDataFile)
	target := data.ReadDataFromCSV(targetDataFile)
	instructions := instruction.NewDecisionClassifier().Parse(instructionFile)

	for i := 0; i < b.N; i++ {
		trainDecisionTreeClassifier(instructions, train, target)
	}
}
func BenchmarkTrainM(b *testing.B) {
	// parser.ReadDataFromCSV()
	trainDataFile, _ := os.Open("./benchmarkdata/x_train.csv")
	targetDataFile, _ := os.Open("./benchmarkdata/y_train.csv")
	instructionFile, _ := os.Open("./benchmarkdata/all.json")
	defer trainDataFile.Close()
	defer targetDataFile.Close()
	defer instructionFile.Close()

	train := data.ReadDataFromCSV(trainDataFile)
	target := data.ReadDataFromCSV(targetDataFile)
	instructions := instruction.Parse(instructionFile)

	for i := 0; i < b.N; i++ {
		trainM(instructions, train, target)
	}
	fmt.Println("trained")
}
