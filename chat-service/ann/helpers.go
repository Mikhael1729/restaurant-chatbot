package ann

import (
	"github.com/jdkato/prose/tokenize"
	"github.com/kljensen/snowball"
)

// normalizeWord stem the given word.
func normalizeWord(word string) string {
	stemmed, err := snowball.Stem(word, "spanish", true)

	if err == nil {
		return stemmed
	}

	return word
}

func tokenizeSentence(sentence string) []string {
	tokenizer := tokenize.NewTreebankWordTokenizer()
	sentenceWords := tokenizer.Tokenize(sentence)

	return sentenceWords
}
