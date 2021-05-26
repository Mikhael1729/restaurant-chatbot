package ann

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Ann struct {
	Dimensions Dimensions
	Inputs     []string
	Outputs    []string
	Parameters Parameters
}

func NewAnn(inputs []string, outputs []string) *Ann {
	dimensions := Dimensions{N0: len(inputs), N1: 20, N2: len(outputs)}
	parameters := Initialize(dimensions.N0, dimensions.N1, dimensions.N2)

	return &Ann{dimensions, inputs, outputs, *parameters}
}

func (ann *Ann) GradientDescent2(x mat.Matrix, y mat.Matrix, alpha float64, iterations int) {
	for i := 0; i < iterations; i++ {
		forward := ann.Parameters.ForwardPropagation(x)
		backward := forward.BackwardPropagation(&ann.Parameters, x, y)
		ann.Parameters.Update(*backward, alpha)

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
