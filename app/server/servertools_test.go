package server

import "testing"

func Test_GetFreeTcpPort(t *testing.T) {
	result, err := GetFreeTcpPort()

	if result < 0 || result > 65535 {
		t.Errorf("Invalid port number: %d", result)
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
