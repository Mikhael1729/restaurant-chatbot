package service

import "time"

type Sender bool

const (
	Customer = true
	Bot      = false
)

type Message struct {
	Id       int
	Text     string
	Sender   bool // If true the sender is the human, otherwise the bot.
	DateTime time.Time
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
