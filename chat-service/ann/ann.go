package ann

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

type Ann struct {
	Dimensions Dimensions
	Inputs     []string
	Outputs    []string
	Parameters Parameters
}

type Dimensions struct {
	N0 int
	N1 int
	N2 int
}

type Parameters struct {
	W1 mat.Matrix
	B1 mat.Matrix
	W2 mat.Matrix
	B2 mat.Matrix
}

type Forward struct {
	Z1 mat.Matrix
	A1 mat.Matrix
	Z2 mat.Matrix
	A2 mat.Matrix
}

type Backward struct {
	DW1 mat.Matrix
	Db1 float64
	DW2 mat.Matrix
	Db2 float64
}

type Examples struct {
	X       mat.Matrix
	Y       mat.Matrix
	Classes mat.Matrix
}

func NewAnn(inputs []string, outputs []string) *Ann {
	n0, n1, n2 := len(inputs), 20, len(outputs)
	dimensions := Dimensions{N0: n0, N1: n1, N2: n2}
	parameters := Parameters{
		W1: mat.NewDense(n1, n0, GenerateRandNorm(n1, n0, 0.01)),
		B1: mat.NewDense(n1, 1, nil),
		W2: mat.NewDense(n2, n1, GenerateRandNorm(n2, n1, 0.01)),
		B2: mat.NewDense(n2, 1, nil),
	}

	return &Ann{
		Dimensions: dimensions,
		Inputs:     inputs,
		Outputs:    outputs,
		Parameters: parameters,
	}
}

func (ann *Ann) GradientDescent(x mat.Matrix, y mat.Matrix, alpha float64, iterations int) {
	for i := 0; i < iterations; i++ {
		forward := ann.ForwardPropagation(x)
		backward := ann.BackwardPropagation(forward, x, y)
		ann.Update(*backward, alpha)

		if i%10 == 0 {
			predictions := getPredictions(forward.A2)
			accuracy := getAccuracy(predictions, y)

			fmt.Printf("Iteration %v: %v", i, accuracy)
			fmt.Println("")
			fmt.Println(predictions)
			fmt.Println("---")
		}
	}
}

func (ann *Ann) ForwardPropagation(X mat.Matrix) *Forward {
	W1, B1, W2, B2 := ann.Parameters.W1, ann.Parameters.B1, ann.Parameters.W2, ann.Parameters.B2

	Z1 := Add(Dot(W1, X), B1)
	A1 := Apply(Relu, Z1)

	Z2 := Add(Dot(W2, A1), B2)
	A2 := Apply(Softmax(Z2), Z2)

	return &Forward{Z1, A1, Z2, A2}
}

func (ann *Ann) BackwardPropagation(forward *Forward, X mat.Matrix, Y mat.Matrix) *Backward {
	_, m := X.Dims()

	oneHotY := oneHot(Y)

	byExamples := func(value float64) float64 {
		return (1.0 / float64(m)) * value
	}

	Z1, A1, A2 := forward.Z1, forward.A1, forward.A2

	dZ2 := Sub(A2, oneHotY)
	dW2 := Apply(byExamples, Dot(dZ2, mat.DenseCopyOf(A1.T())))
	db2 := (1.0 / float64(m)) * mat.Sum(dZ2)

	dZ1 := Multiply(Dot(mat.DenseCopyOf(ann.Parameters.W2.T()), dZ2), Apply(ReluDerivative, Z1))
	dW1 := Apply(byExamples, Dot(dZ1, mat.DenseCopyOf(X.T())))
	db1 := 1.0 / float64(m) * mat.Sum(dZ2)

	return &Backward{dW1, db1, dW2, db2}
}

func (ann *Ann) Update(b Backward, alpha float64) {
	parameters := &ann.Parameters
	W1, B1, W2, B2 := parameters.W1, parameters.B1, parameters.W2, parameters.B2

	timesAlpha := func(value float64) float64 {
		return alpha * value
	}

	W1 = Sub(W1, Apply(timesAlpha, b.DW1))
	B1 = Sub(B1, mat.NewDense(1, 1, []float64{alpha * b.Db1}))

	W2 = Sub(W2, Apply(timesAlpha, b.DW2))
	B2 = Sub(B2, mat.NewDense(1, 1, []float64{alpha * b.Db2}))
}

func getAccuracy(predictions mat.Matrix, Y mat.Matrix) float64 {
	rows, columns := Y.Dims()
	size := rows * columns

	accuracy := float64(Equality(predictions, Y)) / float64(size)

	return accuracy
}

// getPredictions returns a matrix of (m, 1) with the predictions of each sample.
func getPredictions(A2 mat.Matrix) mat.Matrix {
	_, columns := A2.Dims()
	predictionsData := []float64{}

	for j := 0; j < columns; j++ {
		slice := mat.Col(nil, j, A2)

		// Get max value on j-th column.
		max := math.Inf(-1)
		k := 0
		for i, value := range slice {
			intValue := value
			if intValue > max {
				max = intValue
				k = i
			}
		}
		predictionsData = append(predictionsData, float64(k))
	}

	predictions := mat.NewDense(columns, 1, predictionsData)

	return predictions
}

// OneHot2 returns a (m, nL) matrix containing the one-hot arrays for each training example.
func oneHot(Y mat.Matrix) mat.Matrix {
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
