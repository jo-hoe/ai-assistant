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

	id, responseChannel := client.Chat("", []Message{testMessage})
	if len(id) == 0 {
		t.Error("id is empty")
	}
	response := <-responseChannel
	if response.ID != id {
		t.Errorf("expected %s, got %s", id, response.ID)
	}
	if response.Response != testAnswer {
		t.Errorf("expected %s, got %s", testAnswer, response.Response)
	}
	if response.Error != nil {
		t.Errorf("unexpected error %s", response.Error.Error())
	}

	id, responseChannel = client.Chat("", []Message{testMessage})
	if len(id) == 0 {
		t.Error("id is empty")
	}
	response = <-responseChannel
	if response.ID != id {
		t.Errorf("expected %s, got %s", id, response.ID)
	}
	if response.Response != "" {
		t.Errorf("expected %s, got %s", "", response.Response)
	}
	if response.Error == nil {
		t.Error("expected error, got nil")
	}
}
