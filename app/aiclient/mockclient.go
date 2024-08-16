package aiclient

import "time"

const MOCK_CLIENT_TYPE_NAME = "mockclient"

type MockClient struct {
	answers             []string
	delayInMilliseconds int
	err                 error
}

func NewMockClient(answers []string, delayInMilliseconds int, err error) *MockClient {
	return &MockClient{
		answers:             answers,
		delayInMilliseconds: delayInMilliseconds,
		err:                 err,
	}
}

func NewClient(properties map[string]string) (client *AIClient, err error) {
	return nil, nil
}

func (c *MockClient) Chat(messages []Message) (response chan string, err error) {
	if c.err != nil {
		return nil, c.err
	}
	response = make(chan string)
	go c.respond(response)
	return response, err
}

func (c *MockClient) respond(response chan string) {
	for _, answer := range c.answers {
		response <- answer
		time.Sleep(time.Duration(c.delayInMilliseconds))
	}
	close(response)
}
