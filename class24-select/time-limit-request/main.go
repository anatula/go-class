package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// URLs to fetch
	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
	}

	// Fetch each URL with 5 second timeout
	for _, url := range urls {
		fmt.Printf("Fetching: %s\n", url)

		// Create channel for result
		result := make(chan string)

		// Start goroutine to fetch URL
		go func(u string) {
			start := time.Now()
			resp, err := http.Get(u)
			if err != nil {
				result <- fmt.Sprintf("Error: %v", err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				result <- fmt.Sprintf("Error reading body: %v", err)
				return
			}

			// Return success result
			result <- fmt.Sprintf("Success: %s - Status: %d, Size: %d bytes, Time: %v",
				u, resp.StatusCode, len(body), time.Since(start))
		}(url)

		// Wait for result or timeout
		select {
		case res := <-result:
			fmt.Println(res)
		case <-time.After(5 * time.Second):
			fmt.Printf("TIMEOUT: %s took longer than 5 seconds\n", url)
		}

		fmt.Println("---")
	}
}
