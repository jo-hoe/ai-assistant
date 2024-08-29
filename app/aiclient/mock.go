package aiclient

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	MOCK_CLIENT_TYPE_NAME = "mockclient"

	mock_client_err_string              = "errString"
	mock_client_delay_in_ms             = "delayInMilliseconds"
	mock_client_comma_separated_answers = "commaSeparatedAnswers"
)

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
	var answers []string
	var delayInMilliseconds int
	var errString string

	if properties[mock_client_delay_in_ms] == "" {
		delayInMilliseconds = 0
	} else {
		delayInMilliseconds, err = strconv.Atoi(properties[mock_client_delay_in_ms])
		if err != nil {
			return nil, err
		}
	}
	errString = properties[mock_client_err_string]
	if properties[mock_client_comma_separated_answers] == "" {
		answers = []string{}
	} else {
		answers = strings.Split(properties[mock_client_comma_separated_answers], ",")
	}

	return NewMockClient(answers, delayInMilliseconds, errString), nil
}

func (c *MockClient) Ask(query string) (response chan AnswerChunk, err error) {
	if c.errString != "" {
		return nil, errors.New(c.errString)
	}
	response = make(chan AnswerChunk)
	go c.respond(response)
	return response, err
}

func (c *MockClient) respond(response chan AnswerChunk) {
	for i, answerPart := range strings.Split(c.answers[c.count%len(c.answers)], " ") {
		time.Sleep(time.Duration(c.delayInMilliseconds) * time.Millisecond)

		var answer string
		if i == 0 {
			answer = answerPart
		} else {
			answer = " " + answerPart
		}

		answerChunk := AnswerChunk{
			Answer: answer,
		}

		response <- answerChunk
	}
	c.count++
	close(response)
}
