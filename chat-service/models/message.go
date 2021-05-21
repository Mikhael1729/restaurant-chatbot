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
	Sender   bool      `json:"sender,string"` // If true the sender is the human, otherwise the bot.
	DateTime time.Time `json:"dateTime"`
}

func AddMessage(message *Message) {
	message.Id = generateId()
	mockMessages = append(mockMessages, message)
}

func generateId() int {
	lastMessage := mockMessages[len(mockMessages)-1]
	newId := lastMessage.Id + 1

	return newId
}

// FromJson converts a Message in JSON format to a Go Message struct.
func (message *Message) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	decoded := decoder.Decode(message)
	return decoded
}
