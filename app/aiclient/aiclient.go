package aiclient

import (
	"errors"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Answer struct {
	Answer []AnswerChunk
}

type AnswerChunk struct {
	Answer string
	Error  error
}

type AIClients []AIClient

type AIClient interface {
	Chat(messages []Message) (response chan string, err error)
}

func (c *AIClients) GetAnswer(messages []Message) (response chan string, err error) {
	for _, client := range *c {
		response, err = client.Chat(messages)
		if err == nil {
			return response, nil
		}
	}
	return nil, errors.New("all clients failed")
}

func (c *Answer) CompleteAnswer() string {
	stringBuilder := strings.Builder{}
	for _, chunk := range c.Answer {
		stringBuilder.WriteString(chunk.Answer)
	}
	return stringBuilder.String()
}
