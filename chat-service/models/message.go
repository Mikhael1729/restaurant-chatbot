package service

import "time"

type Message struct {
	Text     string
	Sender   bool // If true the sender is the human, otherwise the bot.
	DateTime time.Time
}
