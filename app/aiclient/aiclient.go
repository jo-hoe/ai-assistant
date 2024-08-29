package aiclient

import (
	"errors"
)

type AnswerChunk struct {
	Answer     string
	Err        error
}

func NewAnswerChunk(answer string, err error) *AnswerChunk {
	return &AnswerChunk{
		Answer:     answer,
		Err:        err,
	}
}

type AIClients []AIClient

type AIClient interface {
	Ask(query string) (response chan AnswerChunk, err error)
}

func (c *AIClients) GetAnswer(query string) (response chan AnswerChunk, err error) {
	for _, client := range *c {
		response, err = client.Ask(query)
		if err == nil {
			return response, nil
		}
	}
	return nil, errors.New("all clients failed")
}
