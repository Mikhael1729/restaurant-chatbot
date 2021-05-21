package models

import (
	"encoding/json"
	"io"
	"time"
)

type Sender bool

const (
	Customer = true
	Bot      = false
)

// Message represent a text message between a customer of the restaurant and the bot.
type Message struct {
	Id       int       `json:"id"`
	Text     string    `json:"text"`
	Sender   bool      `json:"sender"` // If true the sender is the human, otherwise the bot.
	DateTime time.Time `json:"dateTime"`
}

type Messages []*Message

func (messages *Messages) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encodedMessages := encoder.Encode(messages)
	return encodedMessages
}

// GetMessages returns the list of stored messages of the chat.
func GetMessages() Messages {
	return mockMessages
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
	{
		Id:       3,
		Text:     "Muy bien, gracias",
		Sender:   Customer,
		DateTime: time.Now(),
	},
	{
		Id:       4,
		Text:     "¿Cómo le puedo ayudar?",
		Sender:   Bot,
		DateTime: time.Now().Add(time.Minute * time.Duration(2)),
	},
}
