package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "https://api.metatrack.ing/health"
	totalRequests := 30

	fmt.Printf("Sending %d requests to %s...\n\n", totalRequests, url)

	var wg sync.WaitGroup
	results := make(chan string, totalRequests)

	// Send requests as fast as possible
	startTime := time.Now()
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go func(reqNum int) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				results <- fmt.Sprintf("Request %d: ERROR - %v", reqNum, err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode == 200 {
				results <- fmt.Sprintf("Request %d: ✅ SUCCESS (200)", reqNum)
			} else if resp.StatusCode == 429 {
				results <- fmt.Sprintf("Request %d: ⛔ RATE LIMITED (429)", reqNum)
			} else {
				results <- fmt.Sprintf("Request %d: Status %d - %s", reqNum, resp.StatusCode, string(body))
			}
		}(i)
		// No delay - send all requests as fast as possible
	}

	wg.Wait()
	close(results)

	elapsed := time.Since(startTime)

	// Print results
	successCount := 0
	rateLimitedCount := 0

	for result := range results {
		fmt.Println(result)
		if contains(result, "SUCCESS") {
			successCount++
		} else if contains(result, "RATE LIMITED") {
			rateLimitedCount++
		}
	}

	fmt.Printf("\n--- Summary ---\n")
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Successful: %d\n", successCount)
	fmt.Printf("Rate limited: %d\n", rateLimitedCount)
	fmt.Printf("Time taken: %v\n", elapsed)
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] != substr &&
		(len(s) > len(substr) && s[:len(substr)] == substr ||
			len(s) >= len(substr) && findSubstr(s, substr))
}

func findSubstr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
