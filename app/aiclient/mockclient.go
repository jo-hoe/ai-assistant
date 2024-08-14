package aiclient

type MockClient struct {
	answers []MockAnswer[string, error]
	count   int
}

type MockAnswer[answer string, err error] struct {
	answer string
	err    error
}

func (c *MockClient) Chat(model string, messages []Message) (responseId string, responseChannel chan *Response) {
	responseChannel = make(chan *Response, 1)

	id := createId()
	go c.respond(id, responseChannel)

	return id, responseChannel
}

func (c *MockClient) respond(responseId string, responseChannel chan *Response) {
	response := NewResponse(
		responseId,
		c.answers[c.count].answer,
		c.answers[c.count].err,
	)
	c.count++

	responseChannel <- response
}

func NewMockClient(answers []MockAnswer[string, error]) *MockClient {
	return &MockClient{
		answers: answers,
		count:   0,
	}
}
