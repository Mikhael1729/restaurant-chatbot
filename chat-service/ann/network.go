package ann

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
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

func (f *Forward) BackwardPropagation(p *Parameters, e *Examples) *Backward {
	_, m := e.X.Dims()

	oneHotY := oneHot(e.Y)

	byExamples := func(value float64) float64 {
		return (1.0 / float64(m)) * value
	}

	dZ2 := Sub(f.A2, oneHotY)
	dW2 := Apply(byExamples, Dot(dZ2, mat.DenseCopyOf(f.A1.T())))
	db2 := (1.0 / float64(m)) * mat.Sum(dZ2)

	dZ1 := Multiply(Dot(mat.DenseCopyOf(p.W2.T()), dZ2), Apply(ReluDerivative, f.Z1))
	dW1 := Apply(byExamples, Dot(dZ1, mat.DenseCopyOf(e.X.T())))
	db1 := 1.0 / float64(m) * mat.Sum(dZ2)

	return &Backward{dW1, db1, dW2, db2}
}

func (p *Parameters) Update(b Backward, alpha float64) *Parameters {
	timesAlpha := func(value float64) float64 {
		return alpha * value
	}

	p.W1 = Sub(p.W1, Apply(timesAlpha, b.DW1))
	p.B1 = Sub(p.B1, mat.NewDense(1, 1, []float64{alpha * b.Db1}))

	p.W2 = Sub(p.W2, Apply(timesAlpha, b.DW2))
	p.B2 = Sub(p.B2, mat.NewDense(1, 1, []float64{alpha * b.Db2}))

	return p
}

func GradientDescent(examples *Examples, alpha float64, iterations int) {
	// Initialize network.
	n0, _ := examples.X.Dims()
	n1 := 20
	n2, _ := examples.Classes.Dims()

	parameters := Initialize(n0, n1, n2)

	for i := 0; i < iterations; i++ {
		forward := parameters.ForwardPropagation(examples.X)
		backward := forward.BackwardPropagation(parameters, examples)
		parameters.Update(*backward, alpha)
		if i%10 == 0 {
			predictions := getPredictions(forward.A2)
			accuracy := getAccuracy(predictions, examples.Y)
			fmt.Printf("Iteration %v: %v", i, accuracy)
			fmt.Println("")
			fmt.Println(predictions)
			fmt.Println("---")
		}
	}
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
