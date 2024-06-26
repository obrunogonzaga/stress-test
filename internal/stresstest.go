package internal

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func StressTest(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	statusCodes := make(map[int]int)
	startTime := time.Now()

	requestsPerGoroutine := totalRequests / concurrency
	remainingRequests := totalRequests % concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int, numRequests int) {
			defer wg.Done()
			client := &http.Client{}
			for j := 0; j < numRequests; j++ {
				resp, err := client.Get(url)
				if err != nil {
					fmt.Printf("Request error: %v\n", err)
					continue
				}
				mu.Lock()
				statusCodes[resp.StatusCode]++
				mu.Unlock()
				resp.Body.Close()
			}
		}(i, requestsPerGoroutine)
	}

	if remainingRequests > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{}
			for j := 0; j < remainingRequests; j++ {
				resp, err := client.Get(url)
				if err != nil {
					fmt.Printf("Request error: %v\n", err)
					continue
				}
				mu.Lock()
				statusCodes[resp.StatusCode]++
				mu.Unlock()
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	totalTime := time.Since(startTime)
	totalStatus200 := statusCodes[http.StatusOK]

	fmt.Printf("Total execution time: %v\n", totalTime)
	fmt.Printf("Total number of requests performed: %d\n", totalRequests)
	fmt.Printf("Number of requests with HTTP status 200: %d\n", totalStatus200)
	fmt.Printf("Distribution of other HTTP status codes:%d\n", len(statusCodes))
	mu.Lock()
	for code, count := range statusCodes {
		if code != http.StatusOK {
			fmt.Printf("Status %d: %d requests\n", code, count)
		}
	}
	mu.Unlock()
}
