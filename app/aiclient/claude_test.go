package aiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClaudeAIClient_Chat(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Check headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("x-api-key") != "test-api-key" {
			t.Errorf("Expected x-api-key: test-api-key, got %s", r.Header.Get("x-api-key"))
		}

		// Decode the request body
		var requestBody map[string]interface{}
		json.NewDecoder(r.Body).Decode(&requestBody)

		// Check the request body
		messages, ok := requestBody["messages"].([]interface{})
		if !ok || len(messages) != 1 {
			t.Errorf("Expected 1 message, got %v", messages)
		}

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"content": []map[string]interface{}{
				{"text": "Hello, I'm Claude!"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create a client that uses the mock server
	client := &ClaudeClient{
		ApiKey:     "test-api-key",
        ApiUrl:     server.URL,
		HttpClient: server.Client(),
	}

	// Test the Chat method
	messages := []Message{
		{Role: "user", Content: "Hello, Claude!"},
	}

	responseChan, err := client.Chat(messages)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Read the response from the channel
	response := <-responseChan

	// Check the response
	expectedResponse := "Hello, I'm Claude!"
	if response != expectedResponse {
		t.Errorf("Expected response %q, got %q", expectedResponse, response)
	}
}
