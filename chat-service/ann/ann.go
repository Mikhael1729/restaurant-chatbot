package ann

import (
	"encoding/json"
	"fmt"
	"github.com/Mikhael1729/restaurant-chatbot/helpers"
	"gonum.org/v1/gonum/mat"
	"math"
)

// AnnModel encodes the data of an ANN in an easy way to convert into JSON format.
type AnnModel struct {
	Dimensions Dimensions `json:"dimensions"`
	Inputs     []string   `json:"inputs"`
	Outputs    []string   `json:"outputs"`
	Parameters map[string][]float64
}

type Ann struct {
	Dimensions Dimensions `json:"dimensions"`
	Inputs     []string   `json:"inputs"`
	Outputs    []string   `json:"outputs"`
	Parameters Parameters `json:"parameters"`
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

func NewAnn(inputs []string, outputs []string) *Ann {
	n0, n1, n2 := len(inputs), 3, len(outputs)
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

func NewAnnModel(ann *Ann) *AnnModel {
	model := &AnnModel{
		Dimensions: ann.Dimensions,
		Inputs:     ann.Inputs,
		Outputs:    ann.Outputs,
		Parameters: map[string][]float64{
			"W1": ann.Parameters.W1.(*mat.Dense).RawMatrix().Data,
			"B1": ann.Parameters.B1.(*mat.Dense).RawMatrix().Data,
			"W2": ann.Parameters.W2.(*mat.Dense).RawMatrix().Data,
			"B2": ann.Parameters.B2.(*mat.Dense).RawMatrix().Data,
		},
	}

	return model
}

func LoadModel(filepath string) (*Ann, error) {
	file, fileError := helpers.GetData(filepath)

	if fileError != nil {
		return nil, fileError
	}

	model := &AnnModel{}
	err := json.Unmarshal(file, model)

	if err != nil {
		return nil, err
	}

	dims := model.Dimensions

	ann := &Ann{
		Inputs:     model.Inputs,
		Dimensions: model.Dimensions,
		Outputs:    model.Outputs,
		Parameters: Parameters{
			W1: mat.NewDense(dims.N1, dims.N0, model.Parameters["W1"]),
			B1: mat.NewDense(dims.N1, 1, model.Parameters["B1"]),
			W2: mat.NewDense(dims.N2, dims.N1, model.Parameters["W2"]),
			B2: mat.NewDense(dims.N2, 1, model.Parameters["B2"]),
		},
	}

	return ann, err
}

func (ann *Ann) SaveModel(filepath string) {
	model := NewAnnModel(ann)

	encoded, err := json.Marshal(model)

	if err != nil {
		panic(err)
	}

	// Create a file, if exists, replace the content.
	file := helpers.CreateFile(filepath)
	defer helpers.CloseFile(file)

	content := string(encoded)
	helpers.WriteFile(file, content)
}

func (ann *Ann) Answer(sentence string) (string, string, float64, int) {
	category, certanty, index := ann.Classify(sentence)
	response := ""
	switch classification := category; classification {
	case "greeting":
		response = "Buenas, espero se encuentre bien"
	case "liked":
		response = "Qué bueno, esperamos atenderle en otra oportunidad"
	case "disliked":
		response = "Sentimos que no haya sido de su agrado. No dude en comunicarse con nosotros sobre cualquier futuro inconveniente"
	case "food,order,hamburger":
		response = "Las opciones para hamburguesa son: Doble queso, Las Rascacielos, BBQ"
	case "food,order,salad":
		response = "Las opciones de ensalada que tenemos son: Ensalada verde, Ensalada de pollo"
	case "food,order,pizza":
		response = "Las opciones de pizza que disponemos son: Pizza Calzone"
	case "food,order,soda":
		response = "Las opciones de soda que disponemos son: Coca-Cola, Pepsi"
	default:
		response = "Su inconveniente lo haremos informar con uno de nuestros asistentes, espere pronta respuesta."
	}

	return response, category, certanty, index
}

func (ann *Ann) Classify(sentence string) (string, float64, int) {
	sentenceInput := ParseSentenceToInput(sentence, ann.Inputs)
	forward := ann.ForwardPropagation(sentenceInput)
	output := forward.A2

	certanty, index := Max(output)
	category := ann.Outputs[index]

	return category, certanty, index
}

func (ann *Ann) GradientDescent(X mat.Matrix, Y mat.Matrix, alpha float64, iterations int) {
	for i := 0; i < iterations; i++ {
		forward := ann.ForwardPropagation(X)
		backward := ann.BackwardPropagation(forward, X, Y)
		ann.Update(*backward, alpha)

		if i%10 == 0 {
			predictions := getPredictions(forward.A2)
			accuracy := getAccuracy(predictions, Y)

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
	A2 := Softmax(Z2)

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
	timesAlpha := func(value float64) float64 {
		return alpha * value
	}

	ann.Parameters.W1 = Sub(ann.Parameters.W1, Apply(timesAlpha, b.DW1))
	ann.Parameters.B1 = Sub(ann.Parameters.B1, mat.NewDense(1, 1, []float64{alpha * b.Db1}))

	ann.Parameters.W2 = Sub(ann.Parameters.W2, Apply(timesAlpha, b.DW2))
	ann.Parameters.B2 = Sub(ann.Parameters.B2, mat.NewDense(1, 1, []float64{alpha * b.Db2}))
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

	maxValue, _ := Max(Y)
	oneHotY := mat.NewDense(rows*columns, int(maxValue)+1, nil)

	for i := 0; i < rows; i++ {
		value := Y.(*mat.Dense).At(i, 0)
		oneHotY.Set(i, int(value), 1)
	}

	// Transpose matrix.
	oneHotY = mat.DenseCopyOf(oneHotY.T())

	return oneHotY
}
