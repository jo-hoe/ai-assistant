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
	return c.answers[c.count].answer, c.answers[c.count].err
}

func NewMockClient(answers []MockAnswer[string, error]) *MockClient {
	return &MockClient{
		answers: answers,
		count:   0,
	}
}
