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

func TestTokenizeSentence(t *testing.T) {
	sentence := "Buenos días, ¿cómo se encuentra?"
	sentenceWords := tokenizeSentence(sentence)
	expected := []string{"Buenos", "días", ",", "¿cómo", "se", "encuentra", "?"}

	for i := 0; i < len(expected); i++ {
		predictedWord := sentenceWords[i]
		expectedWord := expected[i]

		if predictedWord != expectedWord {
			t.Fatalf(`Expected "%v", got "%v" at position %v`, expectedWord, predictedWord, strconv.Itoa(i))
		}
	}
}
