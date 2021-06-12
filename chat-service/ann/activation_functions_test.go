package ann

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestSoftmax(t *testing.T) {
	Z := mat.NewDense(2, 3, []float64{2, -3, 6, 4, 7, 9}) // Pizza, Hola
	resultMatrix := Softmax(Z)

	resultData := resultMatrix.(*mat.Dense).RawMatrix().Data
	expectedData := []float64{0.11920292202211756, 4.53978687024344e-05, 0.04742587317756678, 0.8807970779778824, 0.9999546021312976, 0.9525741268224334}

	perfectMatch, expected, result, i := Verify(expectedData, resultData)
	if !perfectMatch {
		t.Fatalf(`Expected %v, got %v at index %v`, expected, result, i)
	}
}
