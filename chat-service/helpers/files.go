package helpers

import (
	"fmt"
	"os"
)

func CreateFile(path string) *os.File {
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	return file
}

func WriteFile(file *os.File, data interface{}) {
	fmt.Fprintln(file, data)
}

func CloseFile(file *os.File) {
	err := file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
