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

// tokenizeAndSteamSentence takes a text and return the tokenized and stemmed
// words of that given text
func tokenizeAndSteamText(sentence string) []string {
	tokenizer := tokenize.NewTreebankWordTokenizer()
	sentenceWords := tokenizer.Tokenize(sentence)

	// Stem the words in sentenceWords.
	for i := 0; i < len(sentenceWords); i++ {
		stemmedWord := normalizeWord(sentenceWords[i])
		sentenceWords[i] = stemmedWord
	}

	return sentenceWords
}
