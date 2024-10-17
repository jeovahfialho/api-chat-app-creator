package handler

import (
	"net/http"

	"chat-backend/internal/api"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	api.SetupRoutes(router)
	router.ServeHTTP(w, r)
}
