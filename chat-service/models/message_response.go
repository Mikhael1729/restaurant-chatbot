package models

import (
	"encoding/json"
	"io"
)

type MessageResponse struct {
	Message  Message `json:"message"`
	Response Message `json:"response"`
}

func NewMessageResponse(message Message, response Message) *MessageResponse {
	return &MessageResponse{message, response}
}

func (response *MessageResponse) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encodedMessages := encoder.Encode(response)

	return encodedMessages
}
