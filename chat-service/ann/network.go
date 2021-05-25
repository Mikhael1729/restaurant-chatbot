package ann

import (
	"gonum.org/v1/gonum/mat"
)

type Dimensions struct {
	N0 int
	N1 int
	N2 int
}

type Parameters struct {
	W1         mat.Matrix
	B1         mat.Matrix
	W2         mat.Matrix
	B2         mat.Matrix
	Dimensions Dimensions
}

type Forward struct {
	Z1 mat.Matrix
	A1 mat.Matrix
	Z2 mat.Matrix
	A2 mat.Matrix
}

func Initialize(n0 int, n1 int, n2 int) *Parameters {
	W1 := mat.NewDense(n1, n0, GenerateRandNorm(n1, n0, 0.01))
	b1 := mat.NewDense(n1, 1, nil)

	W2 := mat.NewDense(n2, n1, GenerateRandNorm(n2, n1, 0.01))
	b2 := mat.NewDense(n2, 1, nil)

	return &Parameters{W1, b1, W2, b2, Dimensions{n0, n1, n2}}
}

func (p *Parameters) ForwardPropagation(X mat.Matrix) mat.Matrix {
	Z1 := Add(Dot(p.W1, X), p.B1)
	//A1 := Apply(Relu, Z1)

	//Z2 := Add(Dot(p.W2, A1), p.B2)
	//A2 := Apply(Softmax(Z2), Z2)

	//return &Forward{Z1, A1, Z2, A2}
	return Z1
}
