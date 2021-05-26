package handlers

import (
	//"encoding/json"
	"github.com/Mikhael1729/restaurant-chatbot/ann"
	"github.com/Mikhael1729/restaurant-chatbot/models"
	"log"
	"net/http"
)

type Messages struct {
	logger  *log.Logger
	network *ann.Ann
}

func NewMessages(logger *log.Logger) *Messages {
	// Get training data.
	x, y, inputs, outputs := ann.GenerateDevTrainingExamples("./training_data/chats")

	// Create and train the network.
	network := ann.NewAnn(inputs, outputs)
	network.GradientDescent(x, y, 0.10, 500)

	return &Messages{logger, network}
}

func (handler *Messages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	setupCors(rw, r)
	if r.Method == http.MethodGet {
		handler.getMessages(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		handler.addMessage(rw, r)
		return
	}

	// Catch all.
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getMessages process the GET method for the handler
func (m *Messages) getMessages(rw http.ResponseWriter, r *http.Request) {
	messages := models.GetMessages()
	err := messages.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal messages json", http.StatusInternalServerError)
		return
	}
}

// addMessages process the POST method for the handler
func (handler *Messages) addMessage(rw http.ResponseWriter, r *http.Request) {
	message := &models.Message{}

	err := message.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	// Register new message.
	models.AddMessage(message.Text, "", models.Customer)

	// Compute response
	answer, category, _, _ := handler.network.Answer(message.Text)

	// Register message from the bot.
	botMessage := models.AddMessage(answer, category, models.Bot)

	botMessage.ToJson(rw)

	rw.WriteHeader(http.StatusOK)
}

func setupCors(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
