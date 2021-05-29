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

// ReluDerivative
func ReluDerivative(z float64) float64 {
	if z > 0 {
		return 1
	}

	return 0
}

func Softmax(matrix mat.Matrix) mat.Matrix {
	eSum := Sum(Apply(toExp, matrix))
	_, columns := matrix.Dims()

	rows, columns := matrix.Dims()
	outputMatrix := mat.NewDense(rows, columns, nil)

	applyFunc := func(i, j int, value float64) float64 {
		result := math.Exp(matrix.At(i, j)) / eSum.At(0, j)
		return result
	}

	outputMatrix.Apply(applyFunc, matrix)

	return outputMatrix
}

func toExp(value float64) float64 {
	return math.Exp(value)
}

func byESum(sum float64) func(float64) float64 {
	return func(value float64) float64 {
		return math.Exp(value)
	}
}
