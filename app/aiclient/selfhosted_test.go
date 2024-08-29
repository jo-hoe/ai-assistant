package aiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSelfHostedAIClientRespond(t *testing.T) {
	// Create a test server to mock the AI service
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and headers
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("Accept") != "text/event-stream" {
			t.Errorf("Expected Accept: text/event-stream, got %s", r.Header.Get("Accept"))
		}

		// Decode and verify the request body
		var reqBody selfHostedMessageEnvelope
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}
		if reqBody.Model != "test-model" {
			t.Errorf("Expected model 'test-model', got %s", reqBody.Model)
		}
		if len(reqBody.Messages) != 1 || reqBody.Messages[0].Role != defaultRole || reqBody.Messages[0].Content != "test query" {
			t.Errorf("Unexpected messages in request body: %+v", reqBody.Messages)
		}

		// Send mock response
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		responses := []responseData{
			{Type: "provider", Provider: map[string]string{"name": "test-provider"}},
			{Type: "conversation", Conversation: "test-conversation-id"},
			{Type: "content", Content: "Hello"},
			{Type: "content", Content: " World"},
		}
		for _, resp := range responses {
			respJSON, _ := json.Marshal(resp)
			_, _ = w.Write(respJSON)
			_, _ = w.Write([]byte("\n"))
			w.(http.Flusher).Flush()
			time.Sleep(10 * time.Millisecond) // Simulate delay between chunks
		}
	}))
	defer server.Close()

	// Create SelfHostedAIClient with the test server URL
	client := &SelfHostedAIClient{
		Url:    server.URL,
		Model:  "test-model",
		client: server.Client(),
	}

	// Create response channel and call respond
	responseChan := make(chan AnswerChunk)
	go client.respond("test query", responseChan)

	// Collect responses
	var responses []string
	for chunk := range responseChan {
		if chunk.Err != nil {
			t.Errorf("Unexpected error: %v", chunk.Err)
		}
		responses = append(responses, chunk.Answer)
	}

	// Verify responses
	expectedResponses := []string{"Hello", " World"}
	if len(responses) != len(expectedResponses) {
		t.Errorf("Expected %d responses, got %d", len(expectedResponses), len(responses))
	}
	for i, resp := range responses {
		if resp != expectedResponses[i] {
			t.Errorf("Expected response %d to be '%s', got '%s'", i, expectedResponses[i], resp)
		}
	}
}