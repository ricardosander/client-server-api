package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", handleCotacao)

	log.Default().Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Default().Println("Server stopped")
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Received request for /cotacao")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"cotacao": 5.25}`))
}