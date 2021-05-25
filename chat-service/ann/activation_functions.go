package ann

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

// Relu apply ReLU function to the given value z.
func Relu(z float64) float64 {
	if z > 0.00 {
		return z
	} else {
		return 0.00
	}
}

func Softmax(matrix mat.Matrix) func(z float64) float64 {
	return func(z float64) float64 {
		eSum := mat.Sum(Apply(exp, matrix))
		exp := math.Exp(z)

		return exp / eSum
	}
}

func exp(value float64) float64 {
	return math.Exp(value)
}
