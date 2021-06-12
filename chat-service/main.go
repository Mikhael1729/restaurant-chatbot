package main

import (
	"context"
	"fmt"
	"github.com/Mikhael1729/restaurant-chatbot/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	// Load the env variables.
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initializeServer()
}

// initializeServer creates a server with the handlers of the app.
func initializeServer() {
	// Define the logger for the handlers.
	logger := log.New(os.Stdout, "chat-service", log.LstdFlags) // chat-serviceYYY/MM/dd 00:00:00 <message>

	// Create the handlers.
	messagesHandler := handlers.NewMessages(logger, "./saved_models/ann.model13.json")


	// Initialize the router handler.
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger) // Add logger middleware.

	allowedOrigin := os.Getenv("ALLOWED_ORIGINS")
	chiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{allowedOrigin},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
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
		IdleTimeout:  120 * time.Second, // time to keep open resources
		ReadTimeout:  1 * time.Second,   // max duration for reading the request.
		WriteTimeout: 1 * time.Second,   // Max limit to write the response.
	}

	// Not block the code execution by using a go routine.
	go func() {
		err := server.ListenAndServe()
		logger.Fatal(err)
	}()

	// Use the os.Signal to avoid stopping server at the instant until certai signals are received.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until the signal is recived
	logger.Println("Yo have terminated the server", <-sigChan)

	// Close the server. It doesn't accept more requests and finish current work.
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	go func() {
		fmt.Println("antes")
		time.Sleep(time.Second * 2)
		fmt.Println("despues")
		cancel()
	}()
	go func() {
		select {
		case <-timeoutContext.Done():
			fmt.Println(":O")
		}
	}()
	server.Shutdown(timeoutContext)
}
