package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rosariocannavo/api_gateway/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Create a new Gin engine
	router := gin.Default()

	// Create a new instance of RateLimitMiddleware
	rateLimiterMiddleware := middleware.NewRateLimitMiddleware()

	// Use the RateLimitMiddleware in the Gin engine
	router.Use(rateLimiterMiddleware.Handler())
	router.LoadHTMLGlob("../../templates/html/*.html")

	// Define a route that uses the rate-limited middleware
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Number of concurrent requests
	numRequests := 10

	// Channel to synchronize completion of requests
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Make multiple requests concurrently
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()

			// Create a mock HTTP request
			req, err := http.NewRequest("GET", "/protected", nil)
			assert.NoError(t, err)

			// Create a mock HTTP response recorder
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Assert the expected behavior based on the rate limit
			if w.Code == http.StatusOK {
				// Request allowed, add your assertions for allowed scenarios
				assert.Contains(t, w.Body.String(), "success")
			} else {
				// Request rate-limited, add your assertions for rate-limited scenarios
				assert.Equal(t, http.StatusTooManyRequests, w.Code)
			}
		}()
	}

	// Wait for all requests to complete
	wg.Wait()
}
