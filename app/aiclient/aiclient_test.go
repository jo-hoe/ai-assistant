package aiclient

import (
	"testing"
)

func TestAIClients_GetAnswer(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name         string
		c            *AIClients
		args         args
		wantResponse string
		wantErr      bool
	}{
		{
			name: "one client",
			c: &AIClients{
				NewMockClient([]string{"42"}, 0, ""),
			},
			args: args{
				query: "",
			},
			wantResponse: "42",
			wantErr:      false,
		}, {
			name: "multiple clients with a failing one",
			c: &AIClients{
				NewMockClient([]string{""}, 0, "error"),
				NewMockClient([]string{"42"}, 0, ""),
			},
			args: args{
				query: "",
			},
			wantResponse: "42",
			wantErr:      false,
		}, {
			name: "failing client",
			c: &AIClients{
				NewMockClient([]string{""}, 0, "error"),
			},
			args: args{
				query: "",
			},
			wantResponse: "",
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.GetAnswer(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("AIClients.GetAnswer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResponse == nil && tt.wantResponse == "" {
				return
			}

			answer := <-gotResponse
			if answer.Answer != tt.wantResponse {
				t.Errorf("AIClients.GetAnswer() = %v, want %v", answer, tt.wantResponse)
			}
		})
	}
}
