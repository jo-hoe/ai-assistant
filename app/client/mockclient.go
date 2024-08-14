package client

type MockClient struct {
	answers []MockAnswer[string, error]
	count   int
}

type MockAnswer[answer string, err error] struct {
	answer string
	err    error
}

func (c *MockClient) Chat(model string, messages []Message) (string, error) {
	result, err := c.answers[c.count].answer, c.answers[c.count].err
	c.count++
	return result, err
}

func NewMockClient(answers []MockAnswer[string, error]) *MockClient {
	return &MockClient{
		answers: answers,
		count:   0,
	}
}
