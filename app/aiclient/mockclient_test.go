package aiclient

import (
	"errors"
	"strings"
	"testing"
)

func TestMockClient_Chat(t *testing.T) {
	expectedAnswers := []string{"first answer", "second answer"}
	mockClient := NewMockClient(expectedAnswers, 0, nil)

	responseChannel, err := mockClient.Chat(nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	checkAnswer(responseChannel, expectedAnswers, 0, t)

	responseChannel, err = mockClient.Chat(nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	checkAnswer(responseChannel, expectedAnswers, 1, t)
}

func checkAnswer(responseChannel chan string, expectedAnswers []string, index int, t *testing.T) {
	stringBuilder := strings.Builder{}
	for answerPart := range responseChannel {
		stringBuilder.WriteString(answerPart)
	}
	if expectedAnswers[index] != stringBuilder.String() {
		t.Errorf("expected %s, got %s", expectedAnswers[0], stringBuilder.String())
	}
}

func TestMockClient_Chat_Error(t *testing.T) {
	expectedError := errors.New("error")
	mockClient := NewMockClient([]string{}, 0, expectedError)

	responseChannel, err := mockClient.Chat(nil)

	if err == nil {
		t.Errorf("error was nil")
	}

	if responseChannel != nil {
		t.Errorf("response channel was not nil")
	}
}
