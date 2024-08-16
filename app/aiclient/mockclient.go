package aiclient

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const MOCK_CLIENT_TYPE_NAME = "mockclient"

type MockClient struct {
	answers             []string
	delayInMilliseconds int
	err                 error
	count               int
}

func NewMockClient(answers []string, delayInMilliseconds int, err error) *MockClient {
	return &MockClient{
		answers:             answers,
		delayInMilliseconds: delayInMilliseconds,
		err:                 err,
		count:               0,
	}
}

func NewMockClientFromMap(properties map[string]string) (client *MockClient, err error) {
	if properties == nil {
		return nil, errors.New("properties is nil")
	}

	var answers []string
	var delayInMilliseconds int
	var mockErr error

	if properties["delayInMilliseconds"] == "" {
		delayInMilliseconds = 0
	} else {
		delayInMilliseconds, err = strconv.Atoi(properties["delayInMilliseconds"])
		if err != nil {
			return nil, err
		}
	}
	if properties["err"] == "" {
		mockErr = nil
	} else {
		mockErr = errors.New(properties["err"])
	}

	if properties["commaSeparatedAnswers"] == "" {
		answers = []string{}
	} else {
		answers = strings.Split(properties["commaSeparatedAnswers"], ",")
	}

	return NewMockClient(answers, delayInMilliseconds, mockErr), nil
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
	for i, answerPart := range strings.Split(c.answers[c.count], " ") {
		time.Sleep(time.Duration(c.delayInMilliseconds) * time.Millisecond)

		if i == 0 {
			response <- answerPart
		} else {
			response <- " " + answerPart
		}
	}
	c.count++
	close(response)
}
