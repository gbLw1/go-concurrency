package main

import (
	"encoding/json"
	"fmt"
	"go-concurrency/models"
	"net/http"
	"sync"
	"time"
)

func fetchPublicApis(ch chan<- *models.DogApiResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching the API (status code: %s)", resp.Status)
		return
	}

	defer resp.Body.Close()

	var data models.DogApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding response body (status code: %s)\n", resp.Status)
		return
	}

	ch <- &data // register the result in the channel
}

func main() {
	startTime := time.Now()

	ch := make(chan *models.DogApiResponse)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)                   // add a goroutine
		go fetchPublicApis(ch, &wg) // fetch with a goroutine
	}

	// wait all goroutines to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// using the channel to read the data from the goroutine
	i := 0
	for result := range ch {
		i++
		if result != nil {
			fmt.Printf("Fetched %d time(s) -> response: %s\n", i, result.Message)
		}
	}

	fmt.Printf("\nConcurrency operation took: %v", time.Since(startTime))
}
