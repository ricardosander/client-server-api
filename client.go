package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		panic(fmt.Errorf("failed to fetch data: %s", res.Status))
	}
	var data PriceResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

type PriceResponse struct {
	Value string `json:"bid"`
}