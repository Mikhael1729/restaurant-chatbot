package ann

import (
	"strconv"
	"testing"
)

// TestExtractData calls ann.normalizeWord with a word, checking for a valid
// spanish stemmed word.
func TestNormalizeWord(t *testing.T) {
	word := "Avión"
	result := normalizeWord(word)
	expected := "avion"

	if result != expected {
		t.Fatalf(`Expected %v, got %v`, expected, result)
	}
}

// TestTokenizeAndSteamText calls ann.tokenizeAndSteamText with a sentence, checking for a valid
// spanish stemmed and tokenized text.
func TestTokenizeAndSteamText(t *testing.T) {
	sentence := "Buenos días, ¿cómo se encuentra?"
	sentenceWords := tokenizeAndSteamText(sentence)
	expected := []string{"buen", "dias", ",", "¿com", "se", "encuentr", "?"}

	for i := 0; i < len(expected); i++ {
		predictedWord := sentenceWords[i]
		expectedWord := expected[i]

		if predictedWord != expectedWord {
			t.Fatalf(`Expected "%v", got "%v" at position %v`, expectedWord, predictedWord, strconv.Itoa(i))
		}
	}
}

func TestExtractDataSizes(t *testing.T) {
	xSize := 27
	ySize := 27
	inputSize := 33
	outputSize := 7

	data := ExtractData("../training_data/chats")

	x1Size := len(data.X)
	y1Size := len(data.Y)
	input1Size := len(data.InputOptions)
	output1Size := len(data.OutputOptions)

	if xSize != x1Size || ySize != y1Size || inputSize != input1Size || outputSize != output1Size {
		t.Fatalf(`Expected len(X): %v, len(Y): %v, len(InputSize): %v, len(OutputSize): %v, got %v, %v, %v, %v, correspondingly`, xSize, ySize, inputSize, outputSize, x1Size, y1Size, input1Size, output1Size)
	}
}
