package aiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CLAUDE_AI_URL    = "https://api.anthropic.com/v1/messages"
	CLAUDE_TYPE_NAME = "claude"

	claude_api_key_key = "apiKey"
	claude_model_key   = "model"
)

type ClaudeClient struct {
	ApiKey     string
	ApiUrl     string
	Model      string
	HttpClient *http.Client
}

func NewClaudeAIClientFromMap(properties map[string]string) (client *ClaudeClient, err error) {
	var apiKey string
	var model string

	if properties[claude_api_key_key] == "" {
		return nil, fmt.Errorf("missing %s", claude_api_key_key)
	} else {
		apiKey = properties[claude_api_key_key]
	}

	if properties[claude_model_key] == "" {
		model = "claude-2"
	} else {
		model = properties["model"]
	}

	return &ClaudeClient{
		ApiKey:     apiKey,
		ApiUrl:     CLAUDE_AI_URL,
		Model:      model,
		HttpClient: &http.Client{},
	}, nil
}

func (c *ClaudeClient) Chat(messages []Message) (chan string, error) {
	responseChan := make(chan string)

	go func() {
		defer close(responseChan)

		requestBody, err := json.Marshal(map[string]interface{}{
			"model":    c.Model,
			"messages": messages,
		})
		if err != nil {
			responseChan <- fmt.Sprintf("Error marshaling request: %v", err)
			return
		}

		req, err := http.NewRequest("POST", c.ApiUrl, bytes.NewBuffer(requestBody))
		if err != nil {
			responseChan <- fmt.Sprintf("Error creating request: %v", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", c.ApiKey)

		resp, err := c.HttpClient.Do(req)
		if err != nil {
			responseChan <- fmt.Sprintf("Error sending request: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			responseChan <- fmt.Sprintf("Error: Unexpected status code %d", resp.StatusCode)
			return
		}

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			responseChan <- fmt.Sprintf("Error decoding response: %v", err)
			return
		}

		content, ok := result["content"].([]interface{})
		if !ok || len(content) == 0 {
			responseChan <- "Error: Unexpected response format"
			return
		}

		firstContent, ok := content[0].(map[string]interface{})
		if !ok {
			responseChan <- "Error: Unexpected content format"
			return
		}

		text, ok := firstContent["text"].(string)
		if !ok {
			responseChan <- "Error: Unexpected text format"
			return
		}

		responseChan <- text
	}()

	return responseChan, nil
}
