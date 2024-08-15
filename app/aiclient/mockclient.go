package aiclient

const MOCK_CLIENT_TYPE_NAME = "mockclient"

type MockClient struct {
	answers []MockAnswer[string, error]
	count   int
}

type MockAnswer[answer string, err error] struct {
	answer string
	err    error
}

func NewMockClient(answers []MockAnswer[string, error]) *MockClient {
	return &MockClient{
		answers: answers,
		count:   0,
	}
}

func (c *MockClient) Chat(model string, messages []Message) (response string, err error) {
	response, err = c.answers[c.count].answer, c.answers[c.count].err
	c.count++
	return response, err
}
