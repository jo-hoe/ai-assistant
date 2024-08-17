package aiclient

import (
	"errors"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AnswerChunk struct {
	Answer     string
	StopReason string
	Err        error
}

func NewAnswerChunk(answer string, stopReason string, err error) *AnswerChunk {
	return &AnswerChunk{
		Answer:     answer,
		StopReason: stopReason,
		Err:        err,
	}
}

type AIClients []AIClient

type AIClient interface {
	Chat(messages []Message) (response chan AnswerChunk, err error)
}

func (c *AIClients) GetAnswer(messages []Message) (response chan AnswerChunk, err error) {
	for _, client := range *c {
		response, err = client.Chat(messages)
		if err == nil {
			return response, nil
		}
	}
	return nil, errors.New("all clients failed")
}
