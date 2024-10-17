package api

import (
	"encoding/json"
	"log"
	"net/http"

	"chat-backend/internal/claude"
	"chat-backend/internal/models"

	"github.com/gorilla/mux"
)

var currentStep int = 1

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/message", handleMessage).Methods("POST")
	r.HandleFunc("/health", handleHealth).Methods("GET") // Nova rota de health check
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	msg.Content = "User: " + msg.Content + "\nAssistant: "

	response, err := claude.SendMessage(msg.Content)
	if err != nil {
		log.Printf("Error calling Claude: %v", err)
		http.Error(w, "Error processing message", http.StatusInternalServerError)
		return
	}

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
}

// Nova função de health check
func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "OK",
		"message": "Server is running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
