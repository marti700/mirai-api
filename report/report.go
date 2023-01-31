package report

import (
	"bytes"
	"log"
	"text/template"

	"github.com/marti700/mirai/metrics"
	model "github.com/marti700/mirai/models"
	"github.com/marti700/veritas/linearalgebra"
)

// Interface to be implemented by all report structs
type Reporter interface {
	CreateReport(data, target linearalgebra.Matrix, mod model.Model)
	ToString() string
}

// Represents a Regression report, its fields are to be populated by various regression metrics
type RegressionReport struct {
	R2  float64 // r2 metric
	MSE float64 // mean square error
	MAE float64 // mean absolute error
}

// Represents a Classification report, its fields are to be populated by various classification metrics
type ClassificationReport struct {
	ConfusionMatrix []metrics.ConfusionMatrix
}

// Returns an empty instance of a RegressionReport
func NewRegressionReport() RegressionReport {
	return RegressionReport{}
}

// Returns an empty instance of a ClassificationReport
func NewClassificationReport() ClassificationReport {
	return ClassificationReport{}
}

// Creates a regression report based on the provided data and returns it
func (r *RegressionReport) CreateReport(data, target linearalgebra.Matrix, mod model.Model) {
	r.R2 = metrics.RSquared(target, mod.Predict(data))
	r.MSE = metrics.MeanSquareError(target, mod.Predict(data))
}

func (r *RegressionReport) ToString() string {
	templateString := `
Model predictions on the provided test data produced an R squared of {{.R2}} and a Mean Square Error of {{.MSE}}
`
	template := template.Must(template.New("regression").Parse(templateString))
	buff := new(bytes.Buffer)
	err := template.Execute(buff, r)
	if err != nil {
		log.Fatal(err)
	}

	return buff.String()
}

// Creates a classification report based on the provided data and returns it
func (c *ClassificationReport) CreateReport(data, target linearalgebra.Matrix, mod model.Model) {
	cm := metrics.GetConfusionMatrix(target, mod.Predict((data)))
	c.ConfusionMatrix = cm
}

func (r *ClassificationReport) ToString() string {
	templateString := `
Model predictions on the provided test data produced the following result for each classification
{{ $cm := .ConfusionMatrix}}
{{range $cm}}
  - for class x the accuarcy is {{.GetAccuarcy}}, the precision is {{.GetPrecision}} and the Recall value is {{.GetRecall}}
{{end}}
`
	template := template.Must(template.New("classification").Parse(templateString))
	buff := new(bytes.Buffer)
	err := template.Execute(buff, r)
	if err != nil {
		log.Fatal(err)
	}

	return buff.String()
}
