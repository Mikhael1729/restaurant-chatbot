package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {
  fmt.Println("This is my chatbot :\"P\n---")

  trainingData := extractTrainingResponses("./training_data/chats")
  fmt.Println(trainingData)
}

func extractTrainingResponses(filename string) map[string][]string {
  file_data, err := ioutil.ReadFile(filename)

  if err != nil {
    panic(err)
  }

  // Stores the messages grouped by category.
  grouped := make(map[string][]string)

  // Remove null characters.
  text := strings.Replace(string(file_data), "\x00", "", -1)

  // Separate the responses.
  all_responses := strings.Split(text, "#")

  // Group messages into categories.
  for _, response := range all_responses {
    parts := strings.Split(response, "(")

    if len(parts) <= 1 {
      continue
    }

    message := parts[0]
    category := parts[1][:len(parts[1])-1]

    if grouped[category] == nil {
      grouped[category] = []string{message}
    } else {
      grouped[category] = append(grouped[category], message)
    }
  }

  return grouped
}

