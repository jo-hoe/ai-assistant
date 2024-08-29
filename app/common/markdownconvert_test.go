package common

import (
	"reflect"
	"testing"
)

func Test_MarkdownToHTML(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{
				input: "",
			},
			want: "",
		},
		{
			name: "convert",
			args: args{
				input: "# Hello World",
			},
			want: "<h1>Hello World</h1>\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarkdownToHTML(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
