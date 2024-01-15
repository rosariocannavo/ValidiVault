package handlers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rosariocannavo/api_gateway/internal/db"
	"github.com/rosariocannavo/api_gateway/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleLogin_ValidPayload(t *testing.T) {
	err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()
	// Create a new Gin router
	router := gin.Default()

	// Define a route for HandleLogin
	router.POST("/login", handlers.HandleLogin)

	// Create a mock request payload
	payload := map[string]string{
		"Username": "testuser",
		"Password": "testpassword",
	}

	// Convert payload to JSON
	jsonPayload, _ := json.Marshal(payload)

	// Create a mock HTTP request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response code is http.StatusAccepted
	assert.Equal(t, http.StatusAccepted, w.Code)

}

func TestHandleLogin_InvalidPayload(t *testing.T) {
	err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Create a new Gin router
	router := gin.Default()

	// Define a route for HandleLogin
	router.POST("/login", handlers.HandleLogin)

	// Create a mock invalid JSON payload
	invalidPayload := []byte("invalid_payload")

	// Create a mock HTTP request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response code is http.StatusBadRequest
	assert.Equal(t, http.StatusBadRequest, w.Code)

}
