package aiclient

import (
	"errors"
	"testing"
)

func TestMockClient_Chat(t *testing.T) {
	testAnswer := "test"
	testError := errors.New("error")
	testMessage := Message{
		Role:    "user",
		Content: "test",
	}

	client := NewMockClient([]MockAnswer[string, error]{
		{
			answer: testAnswer,
			err:    nil,
		},
		{
			answer: "",
			err:    testError,
		},
	})

	response, err := client.Chat("", []Message{testMessage})
	if response != testAnswer {
		t.Errorf("expected %s, got %s", testAnswer, response)
	}
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	response, err = client.Chat("", []Message{testMessage})
	if response != "" {
		t.Errorf("expected %s, got %s", "", response)
	}
	if err == nil {
		t.Error("expected error, got nil")
	}
}
