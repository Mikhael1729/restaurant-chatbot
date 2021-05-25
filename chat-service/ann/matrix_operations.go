package ann

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

// Apply applies a function to every value of the given matrix.
func Apply(function func(value float64) float64, matrix mat.Matrix) mat.Matrix {
	rows, columns := matrix.Dims()
	outputMatrix := mat.NewDense(rows, columns, nil)

	applyFunc := func(i, j int, value float64) float64 {
		return function(value)
	}

	outputMatrix.Apply(applyFunc, matrix)

	return outputMatrix
}

// GenerateRandomNorm generates a matrix with random numbers with the given dimensions.
func GenerateRandNorm(rows int, columns int, multiplier float64) []float64 {
	size := rows * columns
	data := make([]float64, size)
	for i := range data {
		data[i] = rand.NormFloat64() * multiplier
	}

	return data
}

func GenerateOnes(rows int, columns int) []float64 {
	size := rows * columns
	data := make([]float64, size)
	for i := range data {
		data[i] = 1
	}

	return data
}

// add sum the matrix1 with the matrix2 and returns the result.
func Add(matrix1, matrix2 mat.Matrix) mat.Matrix {
	rows, columns := matrix1.Dims()

	resultMatrix := mat.NewDense(rows, columns, nil)
	resultMatrix.Add(matrix1, matrix2)

	return resultMatrix
}

// dot Multiply mat
func Dot(matrix1, matrix2 mat.Matrix) mat.Matrix {
	rows, _ := matrix1.Dims()
	_, columns := matrix2.Dims()

	product := mat.NewDense(rows, columns, nil)
	product.Product(matrix1, matrix2)

	return product
}