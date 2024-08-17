package aiclient

import (
	"reflect"
	"strings"
	"testing"
)

func TestMockClient_Chat(t *testing.T) {
	expectedAnswers := []string{"first answer", "second answer"}
	mockClient := NewMockClient(expectedAnswers, 0, "")

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

	responseChannel, err = mockClient.Chat(nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	checkAnswer(responseChannel, expectedAnswers, 0, t)
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
	expectedError := "error"
	mockClient := NewMockClient([]string{}, 0, expectedError)

	responseChannel, err := mockClient.Chat(nil)

	if err == nil {
		t.Errorf("error was nil")
	}

	if responseChannel != nil {
		t.Errorf("response channel was not nil")
	}
}

func TestNewMockClientFromMap(t *testing.T) {
	type args struct {
		properties map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantClient *MockClient
		wantErr    bool
	}{
		{
			name: "full properties",
			args: args{
				properties: map[string]string{
					mock_client_delay_in_ms:             "100",
					mock_client_err_string:              "error",
					mock_client_comma_separated_answers: "42,another answer",
				},
			},
			wantClient: &MockClient{
				answers:             []string{"42", "another answer"},
				delayInMilliseconds: 100,
				errString:           "error",
				count:               0,
			},
			wantErr: false,
		}, {
			name: "test defaults",
			args: args{
				properties: map[string]string{},
			},
			wantClient: &MockClient{
				answers:             []string{},
				delayInMilliseconds: 0,
				errString:           "",
				count:               0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := NewMockClientFromMap(tt.args.properties)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMockClientFromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClient, tt.wantClient) {
				t.Errorf("NewMockClientFromMap() = %v, want %v", gotClient, tt.wantClient)
			}
		})
	}
}
