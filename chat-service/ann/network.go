package ann

import (
	"fmt"
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
	dZ1 mat.Matrix
	dW1 mat.Matrix
	db1 mat.Matrix
	dZ2 mat.Matrix
	dW2 mat.Matrix
	db2 mat.Matrix
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

func BackwardPropagation(forward *Forward, parameters *Parameters, examples *Examples) {
	//_, m := examples.X.Dims()

}

// Y(m, 1)
func OneHot(Y mat.Matrix) mat.Dense {
	// Create a (m, nL) array of 0s to store each one hot value.
	rows, columns := Y.Dims()
	fmt.Println(rows, columns)
	dense := Y.(*mat.Dense)
	size := rows * columns

	maxValue := Max(Y)
	oneHotY := *mat.NewDense(size, maxValue, nil)

	for i := 0; i < rows; i++ {
		value := dense.At(i, 0)
		oneHotY.Set(i, int(value), 1)
	}

	oneHotY.T()
	return oneHotY
}

// OneHot2 returns a (m, nL) matrix containing the one-hot arrays for each training example.
//func OneHot2(Y mat.Matrix) mat.Matrix {
//// (m, nL)
//rows, columns := Y.Dims()

//oneHotY := mat.NewDense(rows, Max(Y), nil)

//}

//def one_hot(Y):
//# Create a (m, nL) array of 0s to store each one hot value.
//one_hot_Y = np.zeros((Y.size, Y.max() + 1))

//# Place a one in each position indicated in Y (true labels)
//one_hot_Y[np.arange(Y.size), Y] = 1

//# Use a (nL, m) array instead. Each column is the one hot encoded labels
//one_hot_Y = one_hot_Y.T

//return one_hot_Y
