package ann

import (
	"github.com/jdkato/prose/tokenize"
	"github.com/kljensen/snowball"
	"io/ioutil"
	"strings"
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

// extractData takes the the training file examples to be used to generate
// the input
func ExtractData(dataPath string) *ExtractedData {
	fileData, err := ioutil.ReadFile(dataPath)

	if err != nil {
		panic(err)
	}

	// Clean data by removing null characters.
	cleanText := strings.Replace(string(fileData), "\x00", "", -1)

	// The list of raw data examples.
	examples := strings.Split(cleanText, "#")[1:]

	data := &ExtractedData{
		InputOptions:  map[string]bool{},
		OutputOptions: map[string]bool{},
		X:             [][]string{},
		Y:             []string{},
	}

	for _, example := range examples {
		parts := strings.Split(example, "(")
		category := parts[1][:len(parts[1])-1]
		sentence := parts[0]
		sentenceWords := tokenizeAndSteamText(sentence)

		// Add data of the new example into X and Y.
		data.X = append(data.X, sentenceWords)
		data.Y = append(data.Y, category)

		// Store each new sentence word into inputOptions.
		for _, word := range sentenceWords {
			data.InputOptions[word] = true
		}

		// Store a new category class into outputOptions
		data.OutputOptions[category] = true
	}

	return data
}
