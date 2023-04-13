package report

import (
	"fmt"
	"testing"

	"github.com/marti700/mirai/metrics"
)

func TestClassificationReportToString(t *testing.T) {
	cm := metrics.ConfusionMatrix{
		TP: 10,
		FP: 6,
		FN: 8,
		TN: 9,
	}

	cm2 := metrics.ConfusionMatrix{
		TP: 11,
		FP: 7,
		FN: 9,
		TN: 10,
	}

	tcm := make(map[float64]metrics.ConfusionMatrix)
	tcm[1.0] = cm
	tcm[2.0] = cm2

	cr := ClassificationReport{
		ConfusionMatrices: tcm,
	}

	expectedResult :=
		"\nModel predictions on the provided test data produced the following result for each classification\n \n  - For the class 1: \n      - the accuarcy is 0.5757575757575758, the precision is 0.625 and the Recall value is 0.5555555555555556\n \n  - For the class 2: \n      - the accuarcy is 0.5675675675675675, the precision is 0.6111111111111112 and the Recall value is 0.55\n\n"
	y := cr.ToString()
	fmt.Println(y)

	if cr.ToString() != expectedResult {
		t.Error("A string is expected")
	}
}

func TestRegressionReportToString(t *testing.T) {
	rr := RegressionReport{
		R2:  0.7,
		MSE: 10.4,
	}

	expectedResult := "\nModel predictions on the provided test data produced an R squared of 0.7 and a Mean Square Error of 10.4\n"

	if rr.ToString() != expectedResult {
		t.Error("A string is expected")
	}
}
