package ann

import (
	"github.com/kljensen/snowball"
	"gopkg.in/neurosnap/sentences.v1/english"
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
	tokenizer, err := english.NewSentenceTokenizer(nil)

	if err != nil {
		panic(err)
	}

	tokenizedWords := tokenizer.Tokenize(sentence)
	sentenceWords := []string{}

	for _, token := range tokenizedWords {
		sentenceWords = append(sentenceWords, token.Text)
	}

	return sentenceWords
}
