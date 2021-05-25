package ann

type ExtractedData struct {
	InputOptions  map[string]bool // Unique stemmed sentence words in all examples.
	OutputOptions map[string]bool // Unique outputs (classes)
	X             [][]string      // Stemmed words of all input sentences.
	Y             []string        // The expected output of the examples X.
}
