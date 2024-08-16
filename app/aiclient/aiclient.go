package aiclient

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIClient interface {
	Chat(messages []Message) (response chan string, err error)
}
