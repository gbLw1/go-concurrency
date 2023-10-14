package main

import (
	"encoding/json"
	"fmt"
	"go-concurrency/models"
	"net/http"
	"sync"
	"time"
)

func fetchPublicApis(ch chan<- *models.PublicApiResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://api.publicapis.org/entries"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the APIs")
		return
	}

	defer resp.Body.Close()

	var data models.PublicApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding APIs")
		return
	}

	ch <- &data // register the result in the channel
}

func main() {
	startTime := time.Now()

	ch := make(chan *models.PublicApiResponse)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
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
			fmt.Printf("Fetched %d time(s) -> APIs count: %d\n", i, result.Count)
		}
	}

	fmt.Printf("Concurrency operation took: %v\n\n", time.Since(startTime))
}
