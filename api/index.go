// Package api provides the entry point for all serverless functions.
package api

import (
	"net/http"

	"chat-backend/pkg/api"

	"github.com/gorilla/mux"
)

// Handler represents the entry point for all our serverless functions.
func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	// Configurar rotas
	api.SetupRoutes(router)

	// Adicionar handler para a rota raiz
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	// Servir a requisição
	router.ServeHTTP(w, r)
}
