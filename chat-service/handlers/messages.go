package handlers

import (
	"github.com/Mikhael1729/restaurant-chatbot/models"
	"log"
	"net/http"
)

type Messages struct {
	logger *log.Logger
}

func NewMessages(logger *log.Logger) *Messages {
	return &Messages{logger}
}

func (handler *Messages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handler.getMessages(rw, r)
		return
	}

	// Catch all.
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (m *Messages) getMessages(rw http.ResponseWriter, r *http.Request) {
	messages := models.GetMessages()
	err := messages.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal messages json", http.StatusInternalServerError)
	}
}
