package handlers_test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStressCircuitBreaker(t *testing.T) {
	baseURL := "http://localhost:8080"
	path := "/user/app/product?productId=testProduct"
	concurrency := 50            // Number of concurrent requests
	requestsPerConcurrency := 10 // Number of requests each goroutine will make

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerConcurrency; j++ {
				status := makeRequest(t, baseURL+path)
				assert.True(t, status >= 200 && status <= 300 || status == 429 || status == 401, fmt.Sprintf("Unexpected status code: %d", status)) //500 is not good
				time.Sleep(10 * time.Millisecond)                                                                                                   // Add a small delay between requests
			}
		}()
	}

	wg.Wait()
}

func makeRequest(t *testing.T, url string) int {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
