package main // Note que mudamos para package main

import (
	"net/http"

	"chat-backend/internal/api"

	"github.com/gorilla/mux"
)

// Handler é a função que o Vercel vai chamar
func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	api.SetupRoutes(router)
	router.ServeHTTP(w, r)
}

func main() {
	// Esta função main é necessária para o Vercel, mas não será usada
}
