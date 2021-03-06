package ann

import (
	"errors"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

// Sum returns the sum of the elements on each column of the given matrix.
func Sum(matrix mat.Matrix) mat.Matrix {
	_, columns := matrix.Dims()

	sums := []float64{}
	for j := 0; j < columns; j++ {
		currentColumn := mat.Col(nil, j, matrix)

		sum := 0.0
		for _, value := range currentColumn {
			sum += value
		}

		sums = append(sums, sum)
	}

	sumsMatrix := mat.NewDense(1, columns, sums)

	return sumsMatrix
}

// Max returns the largest value of an integer matrix.
func Max(matrix mat.Matrix) (float64, int) {
	dense := matrix.(*mat.Dense)
	data := dense.RawMatrix().Data

	max := math.Inf(-1)
	index := -1
	for i, value := range data {
		intValue := value
		if intValue > max {
			max = intValue
			index = i
		}
	}

	return max, index
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
	original := data
	for i := 0; i < times-1; i++ {
		data = append(data, original...)
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

// Equality returns the quantity of values that matches on two matrices.
func Equality(matrix1, matrix2 mat.Matrix) int {
	data1 := matrix1.(*mat.Dense).RawMatrix().Data // Data from matrix 1
	data2 := matrix2.(*mat.Dense).RawMatrix().Data // Data from matrix 1

	matchesNumber := 0

	for i := 0; i < len(data1); i++ {
		if data1[i] == data2[i] {
			matchesNumber += 1
		}
	}

	return matchesNumber
}
