package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	var err error 
	db, err = sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return
	}
	defer db.Close()
	if err := initializeDatabase(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", handleCotacao)

	log.Default().Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Default().Println("Server stopped")
}

func initializeDatabase() error {
	createDatabaseQuery := `
	CREATE TABLE IF NOT EXISTS cotacao (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		currency TEXT NOT NULL,
		value REAL NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createDatabaseQuery)
	return err
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	price, err := FindCotacao()
	if err != nil {
		log.Default().Printf("Error fetching cotacao: %v", err)
		http.Error(w, "Failed to fetch cotacao", http.StatusInternalServerError)
		return
	}

	priceValue, ok := (*price)["USDBRL"]
	if !ok {
		log.Default().Println("No data found")
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(priceValue); err != nil {
		log.Default().Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func FindCotacao() (*map[string]Price, error) {
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

	var data map[string]Price
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

type Price struct {
	Value string `json:"bid"`
}
