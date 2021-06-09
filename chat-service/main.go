package main

import (
	"context"
	"github.com/Mikhael1729/restaurant-chatbot/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	initializeServer()
}

// initializeServer creates a server with the handlers of the app.
func initializeServer() {
	// Define the logger for the handlers.
	logger := log.New(os.Stdout, "chat-service", log.LstdFlags) // chat-serviceYYY/MM/dd 00:00:00 <message>

	// Create the handlers.
	messagesHandler := handlers.NewMessages(logger)

	// Initialize the router handler.
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	// Configure cors.
	chiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	
	// Register endpoints.
	chiRouter.Get("/messages", messagesHandler.GetMessages)
	chiRouter.Post("/messages", messagesHandler.AddMessage)

	// Create my own server.
	server := &http.Server{
		Addr:         ":9090",
		Handler:      chiRouter,
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
