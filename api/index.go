package handler

import (
	"log"
	"net/http"
	"os"

	"chat-backend/internal/api"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	api.SetupRoutes(router)

	// Adicione um handler para a rota raiz
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	// Ao invés de iniciar um servidor, usamos o router diretamente
	router.ServeHTTP(w, r)
}

// init function to set up any necessary configurations
func init() {
	// Você pode mover qualquer configuração inicial aqui
	// Por exemplo, carregar variáveis de ambiente
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API configured to run on port %s", port)
}
