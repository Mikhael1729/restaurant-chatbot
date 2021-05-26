package ann

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestSubIntegers(t *testing.T) {
	a := mat.NewDense(2, 2, []float64{3, 3, 3, 3})
	b := mat.NewDense(2, 2, []float64{2, 1, 2, 1})
	c := Sub(a, b)

	resultData := c.(*mat.Dense).RawMatrix().Data
	expectedData := []float64{1, 2, 1, 2}

	perfectMatch, expected, result, i := Verify(expectedData, resultData)
	if !perfectMatch {
		t.Fatalf(`Expected %v, got %v at index %v`, expected, result, i)
	}
}

func TestSubDecimals(t *testing.T) {
	a := mat.NewDense(2, 2, []float64{3.5, 3.5, 3.5, 3.5})
	b := mat.NewDense(2, 2, []float64{2, 1, 2, 1})
	c := Sub(a, b)

	resultData := c.(*mat.Dense).RawMatrix().Data
	expectedData := []float64{1.5, 2.5, 1.5, 2.5}

	perfectMatch, expected, result, i := Verify(expectedData, resultData)
	if !perfectMatch {
		t.Fatalf(`Expected %v, got %v at index %v`, expected, result, i)
	}
}

func TestDotIntegers(t *testing.T) {
	a := mat.NewDense(2, 2, []float64{4, 4, 3, 4})
	b := mat.NewDense(2, 1, []float64{2, 2})
	c := Dot(a, b)

	resultData := c.(*mat.Dense).RawMatrix().Data
	expectedData := []float64{16, 14}

	perfectMatch, expected, result, i := Verify(expectedData, resultData)
	if !perfectMatch {
		t.Fatalf(`Expected %v, got %v at index %v`, expected, result, i)
	}
}

func TestDotDecimals(t *testing.T) {
	a := mat.NewDense(1, 1, []float64{0.2499378781156147})
	b := mat.NewDense(1, 1, []float64{0.2499378781156147})
	c := Dot(a, b)

	resultData := c.(*mat.Dense).RawMatrix().Data
	expectedData := []float64{0.06246894291693586}

	perfectMatch, expected, result, i := Verify(expectedData, resultData)
	if !perfectMatch {
		t.Fatalf(`Expected %v, got %v at index %v`, expected, result, i)
	}
}

func Verify(expectedData []float64, resultData []float64) (bool, float64, float64, int) {
	for i := 0; i < len(resultData); i++ {
		expected := expectedData[i]
		result := resultData[i]

		if expected != result {
			return false, expected, result, i
		}
	}

	return true, 0, 0, 0
}
