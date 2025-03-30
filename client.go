package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "http://localhost:8080/cotacao"
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error:", res.Status)
		panic(fmt.Errorf("failed to fetch data: %s", res.Status))
	}

	var data PriceResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		panic(err)
	}

	err = SaveToFile(data)
	if err != nil {
		fmt.Println("Error saving to file:", err)
		panic(err)
	}

	fmt.Printf("Value %s retrieve from server and saved to file\n", data.Value)
}

type PriceResponse struct {
	Value string `json:"bid"`
}

func SaveToFile(price PriceResponse) error {
	fmt.Printf("Saving to file: %s\n", price.Value)

	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	newLine := fmt.Sprintf("Dólar: %s\n", price.Value)
	_, err = file.WriteString(newLine)
	return err
}