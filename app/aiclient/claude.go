package aiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jo-hoe/ai-assistent/app/common"
)

const (
	CLAUDE_AI_URL              = "https://api.anthropic.com/v1/messages"
	CLAUDE_TYPE_NAME           = "claude"
	CLAUDE_DEFAULT_MODEL       = "claude-3-5-sonnet-20240620"
	claude_default_api_env_key = "CLAUDE_API_KEY"

	claude_api_key_key = "apiKey"
	claude_model_key   = "model"
	claude_api_env_key = "apiEnvKey"
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

	if properties[claude_api_env_key] != "" {
		apiKey = common.GetEnvOrDefault(properties[claude_api_env_key], "")
	}
	if properties[claude_api_key_key] != "" {
		apiKey = properties[claude_api_key_key]
	}

	if properties[claude_model_key] == "" {
		model = CLAUDE_DEFAULT_MODEL
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

func (c *ClaudeClient) Chat(messages []Message) (chan AnswerChunk, error) {
	responseChan := make(chan AnswerChunk)

	go func() {
		defer close(responseChan)

		requestBody, err := json.Marshal(map[string]interface{}{
			"model":    c.Model,
			"messages": messages,
		})
		if err != nil {
			responseChan <- AnswerChunk{
				Answer:     "",
				StopReason: "Error marshaling request",
				Err:        err,
			}
			return
		}

		req, err := http.NewRequest("POST", c.ApiUrl, bytes.NewBuffer(requestBody))
		if err != nil {
			responseChan <- *NewAnswerChunk("", "Error creating request", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", c.ApiKey)

		resp, err := c.HttpClient.Do(req)
		if err != nil {
			responseChan <- *NewAnswerChunk("", "Error sending request", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			message := fmt.Sprintf("Error: Unexpected status code %d", resp.StatusCode)
			responseChan <- *NewAnswerChunk("", message, err)
			return
		}

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			message := fmt.Sprintf("Error decoding response: %v", err)
			responseChan <- *NewAnswerChunk("", message, err)
			return
		}

		content, ok := result["content"].([]interface{})
		if !ok || len(content) == 0 {
			responseChan <- *NewAnswerChunk("", "Error: Unexpected response format", err)
			return
		}

		firstContent, ok := content[0].(map[string]interface{})
		if !ok {
			responseChan <- *NewAnswerChunk("", "Error: Unexpected content format", err)
			return
		}

		text, ok := firstContent["text"].(string)
		if !ok {
			responseChan <- *NewAnswerChunk("", "Error: Unexpected text format", err)
			return
		}

		responseChan <- *NewAnswerChunk(text, "", nil)
	}()

	return responseChan, nil
}
