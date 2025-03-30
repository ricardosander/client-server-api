package main

import (
	"encoding/json"
	"fmt"
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

	cotacao, err := FindCotacao()
	if err != nil {
		log.Default().Printf("Error fetching cotacao: %v", err)
		http.Error(w, "Failed to fetch cotacao", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cotacao); err != nil {
		log.Default().Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func FindCotacao() (interface{}, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", res.Status)
	}
	var data interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}