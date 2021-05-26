package main

import (
	"context"
	"fmt"
	"github.com/Mikhael1729/restaurant-chatbot/ann"
	"github.com/Mikhael1729/restaurant-chatbot/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	useTrainingData()
	//initializeServer()
}

// initializeServer creates a server with the handlers of the app.
func initializeServer() {
	// Define the logger for the handlers.
	logger := log.New(os.Stdout, "chat-service", log.LstdFlags) // chat-serviceYYY/MM/dd 00:00:00 <message>

	// Create the handlers.
	messagesHandler := handlers.NewMessages(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/messages", messagesHandler)

	// Create my own server.
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Not block the code execution by using a go routine.
	go func() {
		err := server.ListenAndServe()
		logger.Fatal(err)
	}()

	// Use the os.Signal to avoid stopping server at the instant until a certain
	// signals are executed.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until the signal is recived
	sig := <-sigChan
	logger.Println("Yo have terminated the server", sig)

	// Close the server. It doesn't accept more requests and finish current work.
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

// useTrainingData extracts the training data and shows it in the console.
func useTrainingData() {
	// Get training data.
	x, y, inputs, outputs := ann.GenerateDevTrainingExamples("./training_data/chats")

	// Create and train the network.
	network := ann.NewAnn(inputs, outputs)
	network.GradientDescent(x, y, 0.10, 500)

	// Make prediction.
	sentence := "Buen día"
	category, certanty, _ := network.Classify(sentence)
	fmt.Printf("Frase: %v\n", sentence)
	fmt.Printf("category: %v\ncertanty: %v\n", category, certanty)
}
