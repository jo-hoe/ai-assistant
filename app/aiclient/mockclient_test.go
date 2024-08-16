package aiclient

import (
	"errors"
	"testing"
)

func TestMockClient_Chat(t *testing.T) {
	expectedAnswers := []string{"answer1", "answer2"}
	mockClient := NewMockClient(expectedAnswers, 0, nil)

	responseChannel, err := mockClient.Chat(nil)

	if err != nil {
		t.Errorf(err.Error())
	}

	for _, expectedAnswer := range expectedAnswers {
		answer := <-responseChannel
		if answer != expectedAnswer {
			t.Errorf("Expected %s, got %s", expectedAnswer, answer)
		}
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
