package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:8080/cotacao"
	req, err := http.NewRequest("GET", url, nil)
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

	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	newLine := fmt.Sprintf("Dólar: %s\n", price.Value)
	_, err = file.WriteString(newLine)
	return err
}