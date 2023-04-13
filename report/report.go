package report

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/marti700/mirai/metrics"
	model "github.com/marti700/mirai/models"
	"github.com/marti700/veritas/linearalgebra"

	goemail "gopkg.in/gomail.v2"
)

// Interface to be implemented by all report structs
type Reporter interface {
	CreateReport(actual, predicted linearalgebra.Matrix, mod model.Model)
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
	ConfusionMatrices map[float64]metrics.ConfusionMatrix
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
func (r *RegressionReport) CreateReport(actual, predicted linearalgebra.Matrix, mod model.Model) {
	r.R2 = metrics.RSquared(actual, predicted)
	r.MSE = metrics.MeanSquareError(actual, predicted)
}

// ToString implementation of the Reporter interface
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
func (c *ClassificationReport) CreateReport(actual, predicted linearalgebra.Matrix, mod model.Model) {
	cm := metrics.BuildConfusionMatrices(actual, predicted)
	c.ConfusionMatrices = cm
}

// ToString implementation of the Reporter interface
func (r *ClassificationReport) ToString() string {
	// golang templates can iterate a map just if it's keys are strings. classificationReport#ConfusionMatrix is
	// a map wich key type is a float64 so, a new map with string key type is needed to comply with the golang text
	// template constraints
	data := make(map[string]metrics.ConfusionMatrix)
	for k, v := range r.ConfusionMatrices {
		data[strconv.FormatFloat(k, 'f', -1, 64)] = v
	}

	templateString := `
Model predictions on the provided test data produced the following result for each classification
{{range $key, $value := .}}
  - For the class {{ $key }}:
      - the accuarcy is {{$value.GetAccuarcy}}, the precision is {{$value.GetPrecision}} and the Recall value is {{$value.GetRecall}}
{{end}}
`
	template := template.Must(template.New("classification").Parse(templateString))
	buff := new(bytes.Buffer)
	err := template.Execute(buff, data)
	if err != nil {
		log.Fatal(err)
	}

	return buff.String()
}

func SendReportByEmail(email, attachmentPath string) error {

	emailSender := os.Getenv("SENDER_EMAIL")
	msg := goemail.NewMessage()
	msg.SetHeader("From", emailSender)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Mirai reports")
	msg.SetBody("text/html", "This is a test mail")
	msg.Attach(attachmentPath)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	n := goemail.NewDialer(os.Getenv("SMTP"), port, emailSender, os.Getenv("SENDER_EMAIL_PASSWORD"))

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
