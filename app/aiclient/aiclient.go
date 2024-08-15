package aiclient

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIClient interface {
	Chat(model string, messages []Message) (response string, err error)
}
