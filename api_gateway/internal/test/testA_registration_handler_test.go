// registration_handler_test.go
package handlers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rosariocannavo/api_gateway/internal/db"
	"github.com/rosariocannavo/api_gateway/internal/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHandleRegistration_ValidPayload(t *testing.T) {

	err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Create a new Gin router
	router := gin.Default()

	// Define a route for HandleRegistration
	router.POST("/register", handlers.HandleRegistration)

	// Create a mock request payload
	payload := map[string]string{
		"Username":        "testuser",
		"Password":        "testpassword",
		"MetamaskAddress": "0x58ad8fEA5d85EDD13C05dC116198801Ff53679B2",
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error converting map to JSON:", err)
	}

	// Create a mock HTTP request
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response code is http.StatusOK
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestHandleRegistration_InvalidPayload(t *testing.T) {

	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Create a new Gin router
	router := gin.Default()

	// Define a route for HandleRegistration
	router.POST("/register", handlers.HandleRegistration)

	// Create a mock request payload with missing fields
	payload := map[string]string{
		// Missing "Password" field
		"Username":        "testuser",
		"MetamaskAddress": "0x58ad8fEA5d85EDD13C05dC116198801Ff53679B2",
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error converting map to JSON:", err)
	}

	// Create a mock HTTP request
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response code is http.StatusBadRequest for an invalid payload
	assert.Equal(t, http.StatusForbidden, w.Code)
}
