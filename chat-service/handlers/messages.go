package handlers

import (
	"github.com/Mikhael1729/restaurant-chatbot/ann"
	"github.com/Mikhael1729/restaurant-chatbot/models"
	"log"
	"net/http"
)

type Messages struct {
	logger  *log.Logger
	network *ann.Ann
}

func NewMessages(logger *log.Logger, modelPath string) *Messages {
	// Create and train the network.
	network, err := ann.LoadModel(modelPath)

	if err != nil {
		// Get training data.
		x, y, inputs, outputs := ann.GenerateDevTrainingExamples("./training_data/chats")
		network = ann.NewAnn(inputs, outputs)
		network.GradientDescent(x, y, 0.9, 10000)

		// Save the model.
		network.SaveModel(modelPath)
		logger.Printf("A new model has been saved on %v\n", modelPath)
	}

	logger.Printf("The ANN model has been loaded succesfully from %v\n", modelPath)

	return &Messages{logger, network}
}

// MetMessages process the GET method for the handler
func (m *Messages) GetMessages(rw http.ResponseWriter, r *http.Request) {
	messages := models.GetMessages()
	err := messages.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal messages json", http.StatusInternalServerError)
		return
	}
}

// AddMessages process the POST method for the handler
func (handler *Messages) AddMessage(rw http.ResponseWriter, r *http.Request) {
	message := &models.Message{}

	err := message.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	// Register new message.
	message = models.AddMessage(message.Text, "", models.Customer)

	// Compute response
	answer, category, _, _, _ := handler.network.Answer(message.Text)

	// Register message from the bot.
	botMessage := models.AddMessage(answer, category, models.Bot)

	// Send response.
	messageResponse := models.NewMessageResponse(*message, *botMessage)
	messageResponse.ToJson(rw)
}
