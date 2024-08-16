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
	errString           string
	count               int
}

func NewMockClient(answers []string, delayInMilliseconds int, errString string) *MockClient {
	return &MockClient{
		answers:             answers,
		delayInMilliseconds: delayInMilliseconds,
		errString:           errString,
		count:               0,
	}
}

func NewMockClientFromMap(properties map[string]string) (client *MockClient, err error) {
	if properties == nil {
		return nil, errors.New("properties is nil")
	}

	var answers []string
	var delayInMilliseconds int
	var errString string

	if properties["delayInMilliseconds"] == "" {
		delayInMilliseconds = 0
	} else {
		delayInMilliseconds, err = strconv.Atoi(properties["delayInMilliseconds"])
		if err != nil {
			return nil, err
		}
	}
	errString = properties["err"]
	if properties["commaSeparatedAnswers"] == "" {
		answers = []string{}
	} else {
		answers = strings.Split(properties["commaSeparatedAnswers"], ",")
	}

	return NewMockClient(answers, delayInMilliseconds, errString), nil
}

func (c *MockClient) Chat(messages []Message) (response chan string, err error) {
	if c.errString != "" {
		return nil, errors.New(c.errString)
	}
	response = make(chan string)
	go c.respond(response)
	return response, err
}

func (c *MockClient) respond(response chan string) {
	for i, answerPart := range strings.Split(c.answers[c.count], " ") {
		time.Sleep(time.Duration(c.delayInMilliseconds*1000) * time.Millisecond)

		if i == 0 {
			response <- answerPart
		} else {
			response <- " " + answerPart
		}
	}
	c.count++
	close(response)
}
