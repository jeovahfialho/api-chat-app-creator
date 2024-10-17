package handler

import (
	"log"
	"net/http"
	"os"

	"chat-backend/internal/api"

	"github.com/gorilla/mux"
)

func Handler() {
	r := mux.NewRouter()
	api.SetupRoutes(r)

	// Adicione um handler para a rota raiz
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
