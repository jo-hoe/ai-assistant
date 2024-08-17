package aiclient

import "errors"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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
