package aiclient

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type SelfHostedAIClient struct {
	url    string
	model  string
	client *http.Client
}

type selfHostedMessageEnvelope struct {
	Model    string              `json:"model"`
	Messages []selfHostedMessage `json:"messages"`
}

type responseData struct {
	Type         string      `json:"type"`
	Provider     interface{} `json:"provider,omitempty"`
	Conversation string      `json:"conversation,omitempty"`
	Content      string      `json:"content,omitempty"`
}

const (
	defaultRole = "assistant"
)

type selfHostedMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewSelfHostedAIClient(url string, model string) *SelfHostedAIClient {
	return &SelfHostedAIClient{
		url:   url,
		model: model,
	}
}

func (c *SelfHostedAIClient) Ask(query string) (response chan AnswerChunk, err error) {
	response = make(chan AnswerChunk)
	go c.respond(query, response)
	return response, err
}

func (c *SelfHostedAIClient) respond(query string, response chan AnswerChunk) {
	message := selfHostedMessageEnvelope{
		Model: c.model,
		Messages: []selfHostedMessage{
			{
				Role:    defaultRole,
				Content: query,
			},
		},
	}

	// Convert message to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		response <- *NewAnswerChunk("", err)
		close(response)
		return
	}

	// Create a new request
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonData))
	if err != nil {
		response <- *NewAnswerChunk("", err)
		close(response)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	// Send the request
	resp, err := c.client.Do(req)
	if err != nil {
		response <- *NewAnswerChunk("", err)
		close(response)
		return
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		response <- *NewAnswerChunk("", err)
		close(response)
		return
	}

	// Read the response body
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			response <- *NewAnswerChunk("", err)
			close(response)
			return
		}

		// Process the received event
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var data responseData
		err = json.Unmarshal([]byte(line), &data)
		if err != nil {
			response <- *NewAnswerChunk("", err)
			close(response)
			return
		}

		switch data.Type {
		case "provider":
			log.Printf("received answer from provider: %+v\n", data.Provider)
		case "conversation":
			log.Printf("conversion ID: %+v\n", data.Conversation)
		case "content":
			log.Printf("content received: %+v\n", data.Content)
			response <- *NewAnswerChunk(data.Content, nil)
		default:
			log.Printf("unknown type: %s\n", data.Type)
		}
	}

	close(response)
}
