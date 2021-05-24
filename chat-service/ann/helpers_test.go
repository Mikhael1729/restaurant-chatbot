package ann

import (
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
