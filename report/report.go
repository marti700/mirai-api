package report

import (
	"github.com/marti700/mirai/metrics"
	model "github.com/marti700/mirai/models"
	"github.com/marti700/veritas/linearalgebra"
)

// Interface to be implemented by all report structs
type Reporter interface {
	CreateReport(data, target linearalgebra.Matrix, mod model.Model)
}

// Represents a Regression report, its fields are to be populated by various regression metrics
type RegressionReport struct {
	r2  float64 // r2 metric
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
	r.r2 = metrics.RSquared(target, mod.Predict(data))
	r.MSE = metrics.MeanSquareError(target, mod.Predict(data))
}

// Creates a classification report based on the provided data and returns it
func (c *ClassificationReport) CreateReport(data, target linearalgebra.Matrix, mod model.Model) {
	cm := metrics.GetConfusionMatrix(target, mod.Predict((data)))
	c.ConfusionMatrix = cm
}
