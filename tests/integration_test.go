package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
)

func TestHealthzHandlerIntegration(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(handlers.HealthzHandler))
	defer server.Close()

	// Log the server URL
	t.Logf("Test server URL: %s", server.URL)

	// Perform a GET request to the /healthz endpoint
	response, err := http.Get(server.URL + "/healthz")
	if err != nil {
		t.Fatalf("Error sending GET request: %v", err)
	}
	defer response.Body.Close()

	// Log the response status code
	t.Logf("Response status code: %d", response.StatusCode)

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}
