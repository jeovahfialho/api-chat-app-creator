package main

import (
	"log"
	"net/http"

	"chat-backend/internal/api"
	"chat-backend/pkg/config"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	r := mux.NewRouter()
	api.SetupRoutes(r)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
