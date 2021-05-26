package ann

import (
	"errors"
	//"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

// Max returns the largest value of an integer matrix.
func Max(matrix mat.Matrix) int {
	dense := matrix.(*mat.Dense)
	data := dense.RawMatrix().Data

	max := int(math.Inf(-1))
	for _, value := range data {
		intValue := int(value)
		if intValue > max {
			max = intValue
		}
	}

	return max
}

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
	rows2, columns2 := matrix2.Dims()

	resultMatrix := mat.NewDense(rows, columns, nil)

	if rows*columns > rows2*columns2 {
		resultMatrix.Add(matrix1, Broadcast(matrix2, rows, columns))
	} else if rows*columns == rows2*columns2 {
		resultMatrix.Add(matrix1, matrix2)
	} else {
		resultMatrix.Add(Broadcast(matrix1, rows2, columns2), matrix2)
	}

	return resultMatrix
}

func Sub(matrix1, matrix2 mat.Matrix) mat.Matrix {
	rows, columns := matrix1.Dims()
	rows2, columns2 := matrix2.Dims()

	resultMatrix := mat.NewDense(rows, columns, nil)
	if rows*columns > rows2*columns2 {
		resultMatrix.Sub(matrix1, Broadcast(matrix2, rows, columns))
	} else if rows*columns == rows2*columns2 {
		resultMatrix.Sub(matrix1, matrix2)
	} else {
		resultMatrix.Sub(Broadcast(matrix1, rows2, columns2), matrix2)
	}

	return resultMatrix
}

func Broadcast(matrix mat.Matrix, rows int, columns int) mat.Matrix {
	dense := matrix.(*mat.Dense)
	prevRows, prevColumns := dense.Dims()

	newDataSize := rows * columns
	previousDataSize := prevRows * prevColumns

	if newDataSize%previousDataSize != 0 {
		panic(errors.New("The matrix couldn't be broadcasted"))
	}

	// Generate the apropiate amount of data for the new size
	times := newDataSize / previousDataSize
	data := dense.RawMatrix().Data
	for i := 0; i < times-1; i++ {
		data = append(data, data...)
	}

	return mat.NewDense(rows, columns, data)
}

// Dot Multiply mat
func Dot(matrix1, matrix2 mat.Matrix) mat.Matrix {
	rows, _ := matrix1.Dims()
	_, columns := matrix2.Dims()

	product := mat.NewDense(rows, columns, nil)
	product.Product(matrix1, matrix2)

	return product
}

// Multiply
func Multiply(matrix1, matrix2 mat.Matrix) mat.Matrix {
	rows, columns := matrix1.Dims()

	result := mat.NewDense(rows, columns, nil)
	result.MulElem(matrix1, matrix2)

	return result
}
