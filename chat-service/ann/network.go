package ann

import (
	//"fmt"
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

type Backward struct {
	dW1 mat.Matrix
	db1 float64
	dW2 mat.Matrix
	db2 float64
}

type Examples struct {
	X mat.Matrix
	Y mat.Matrix
}

func Initialize(n0 int, n1 int, n2 int) *Parameters {
	W1 := mat.NewDense(n1, n0, GenerateRandNorm(n1, n0, 0.01))
	b1 := mat.NewDense(n1, 1, nil)

	W2 := mat.NewDense(n2, n1, GenerateRandNorm(n2, n1, 0.01))
	b2 := mat.NewDense(n2, 1, nil)

	return &Parameters{W1, b1, W2, b2, Dimensions{n0, n1, n2}}
}

func (p *Parameters) ForwardPropagation(X mat.Matrix) *Forward {
	Z1 := Add(Dot(p.W1, X), p.B1)
	A1 := Apply(Relu, Z1)

	Z2 := Add(Dot(p.W2, A1), p.B2)
	A2 := Apply(Softmax(Z2), Z2)

	return &Forward{Z1, A1, Z2, A2}
}

func BackwardPropagation(f *Forward, p *Parameters, e *Examples) *Backward {
	_, m := e.X.Dims()

	oneHotY := OneHot(e.Y)

	byExamples := func(value float64) float64 { return float64(1/m) * value }

	dZ2 := Sub(f.A2, oneHotY)
	dW2 := Apply(byExamples, Dot(dZ2, f.A2))
	db2 := float64(1/m) * mat.Sum(dZ2)

	dZ1 := Multiply(Dot(p.W2.T(), dZ2), Apply(ReluDerivative, f.Z1))
	dW1 := Apply(byExamples, Dot(dZ1, e.X.T()))
	db1 := float64(1/m) * mat.Sum(dZ2)

	return &Backward{dW1, db1, dW2, db2}
}

// OneHot2 returns a (m, nL) matrix containing the one-hot arrays for each training example.
func OneHot(Y mat.Matrix) mat.Matrix {
	rows, columns := Y.Dims()

	maxValue := Max(Y)
	oneHotY := mat.NewDense(rows*columns, maxValue+1, nil)

	for i := 0; i < rows; i++ {
		value := Y.(*mat.Dense).At(i, 0)
		oneHotY.Set(i, int(value), 1)
	}

	// Transpose matrix.
	oneHotY = mat.DenseCopyOf(oneHotY.T())

	return oneHotY
}
