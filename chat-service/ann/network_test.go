package ann

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestGetPredictions(t *testing.T) {
	matrix := mat.NewDense(3, 2, []float64{2, 3, 3, 2, 1, 2})
	predictions := getPredictions(matrix)

	expectedData := []float64{1, 0}
	resultData := predictions.(*mat.Dense).RawMatrix().Data

	perfectMatch, expected, result, i := Verify(expectedData, resultData)
	if !perfectMatch {
		t.Fatalf(`Expected %v, got %v at index %v`, expected, result, i)
	}
}
