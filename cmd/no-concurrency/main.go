package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-concurrency/models"
	"net/http"
	"time"
)

func fetchPublicApis() (*models.PublicApiResponse, error) {
	var data models.PublicApiResponse

	url := "https://api.publicapis.org/entries"
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("error fetching the APIs")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.New("error decoding APIs")
	}

	return &data, nil
}

func main() {
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		data, err := fetchPublicApis()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
		}

		fmt.Printf("Fetched %d time(s) -> APIs count: %d\n", i+1, data.Count)
	}

	fmt.Printf("NO concurrency operation took: %v\n\n", time.Since(startTime))
}
