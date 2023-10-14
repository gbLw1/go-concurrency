package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-concurrency/models"
	"net/http"
	"time"
)

func fetchPublicApis() (*models.DogApiResponse, error) {
	var data models.DogApiResponse

	url := "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("error fetching the API -> status code: " + resp.Status)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.New("error decoding response body -> status code: " + resp.Status)
	}

	return &data, nil
}

func main() {
	startTime := time.Now()

	for i := 0; i < 20; i++ {
		data, err := fetchPublicApis()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Fetched %d time(s) -> response: %s\n", i+1, data.Message)
	}

	fmt.Printf("\nNO concurrency operation took: %v", time.Since(startTime))
}
