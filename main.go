package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(2)
	}
}

func main() {
	fmt.Println("This is my chatbot :\"P\n---\n")

	file_data, err := ioutil.ReadFile("./training_data/chats")

	if err != nil {
		panic(err)
	}

	text := string(file_data)
	all_responses := strings.Split(text, "#")

	for _, message := range all_responses {
		fmt.Println(message)
	}
}

type CategoryMessages struct {
	category string
	messages []string
}
