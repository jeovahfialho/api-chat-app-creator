package main

import (
	"fmt"
	"net/http"

	"chat-backend/internal/api"

	"github.com/gorilla/mux"
)

// Handler é a função que o Vercel vai chamar
func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	api.SetupRoutes(router)

	// Adiciona o handler para a rota raiz
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	// Serve a requisição usando o router
	router.ServeHTTP(w, r)
}

// Mantemos a função main para testes locais
func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
