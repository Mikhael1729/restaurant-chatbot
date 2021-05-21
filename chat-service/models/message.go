package models

import (
	"encoding/json"
	"io"
	"time"
)

type Sender int

const (
	Customer = 0
	Bot      = 1
)

// Message represent a text message between a customer of the restaurant and the bot.
type Message struct {
	Id       int       `json:"id"`
	Text     string    `json:"text"`
	Sender   Sender    `json:"sender"` // If true the sender is the human, otherwise the bot.
	DateTime time.Time `json:"dateTime"`
}

type Messages []*Message

// GetMessages returns the list of stored messages of the chat.
func GetMessages() Messages {
	return mockMessages
}

// AddMessage adds a new Message to the database.
func AddMessage(message *Message) {
	message.Id = generateId()
	mockMessages = append(mockMessages, message)
}

// ToJson converts a Messages into JSON format.
func (messages *Messages) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encodedMessages := encoder.Encode(messages)
	return encodedMessages
}

// FromJson converts a Message in JSON format to a Go Message struct.
func (message *Message) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	decoded := decoder.Decode(message)
	return decoded
}

// generateId generate a new Id for a new message.
func generateId() int {
	lastMessage := mockMessages[len(mockMessages)-1]
	newId := lastMessage.Id + 1

	return newId
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
