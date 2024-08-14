package aiclient

import (
	"github.com/google/uuid"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	ID       string `json:"id"`
	Response string `json:"response"`
	Error    error  `json:"error"`
}

func NewResponse(id string, response string, err error) *Response {
	return &Response{
		ID:       id,
		Response: response,
		Error:    err,
	}
}

type AIClient interface {
	Chat(model string, messages []Message) (responseId string, responseChannel chan *Response)
}

func createId() string {
	id := uuid.New()
	return id.String()
}
