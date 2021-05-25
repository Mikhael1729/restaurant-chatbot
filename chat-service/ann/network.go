package ann

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

type Ann struct {
	W1 mat.Dense
	B1 mat.Dense
	W2 mat.Dense
	B2 mat.Dense
}

func NewAnn(n0 int, n1 int, n2 int) *Ann {
	W1 := mat.NewDense(n1, n0, generateRandNorm(n1, n0, 0.01))
	b1 := mat.NewDense(n1, 1, nil)

	W2 := mat.NewDense(n2, n1, generateRandNorm(n2, n1, 0.01))
	b2 := mat.NewDense(n2, 1, nil)

	return &Ann{*W1, *b1, *W2, *b2}
}

func generateRandNorm(r int, c int, multiplier float64) []float64 {
	size := r * c
	data := make([]float64, size)
	for i := range data {
		data[i] = rand.NormFloat64() * multiplier
	}

	return data
}
