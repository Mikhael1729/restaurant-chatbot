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
