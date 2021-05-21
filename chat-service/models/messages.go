package models

import (
	"encoding/json"
	"io"
	"time"
)

type Messages []*Message

// GetMessages returns the list of stored messages of the chat.
func GetMessages() Messages {
	return mockMessages
}

// ToJson converts a Messages into JSON format.
func (messages *Messages) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encodedMessages := encoder.Encode(messages)
	return encodedMessages
}

var mockMessages = []*Message{
	{
		Id:       1,
		Text:     "Hola",
		Sender:   Customer,
		DateTime: time.Now(),
	},
	{
		Id:       2,
		Text:     "Buenos días, ¿cómo se encuentra?",
		Sender:   Bot,
		DateTime: time.Now().Add(time.Minute * time.Duration(2)),
	},
}
