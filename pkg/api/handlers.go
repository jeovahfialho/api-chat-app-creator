package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"chat-backend/internal/claude"
	"chat-backend/internal/models"

	"github.com/gorilla/mux"
)

var currentStep int = 1

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/message", handleMessage).Methods("POST")
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to /api/message")

	// Read and log the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}
	log.Printf("Request body: %s", string(body))

	var msg models.Message
	if err := json.Unmarshal(body, &msg); err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded message: %+v", msg)

	if msg.Content == "" {
		log.Println("Error: Empty message content")
		http.Error(w, "Message content cannot be empty", http.StatusBadRequest)
		return
	}

	msg.Content = fmt.Sprintf("User: %s\nAssistant: ", msg.Content)

	log.Printf("Sending message to Claude: %s", msg.Content)
	response, err := claude.SendMessage(msg.Content)
	if err != nil {
		log.Printf("Error calling Claude: %v", err)
		http.Error(w, "Error processing message", http.StatusInternalServerError)
		return
	}

	log.Printf("Received response from Claude: %s", response)

	currentStep++

	responseMsg := models.Message{
		Content: response,
		Step:    currentStep,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseMsg); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	log.Println("Response sent successfully")
}
