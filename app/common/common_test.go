package common

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	envKey := "TEST_ENV_KEY"
	envValue := "TEST_ENV_VALUE"

	os.Setenv(envKey, envValue)
	defer os.Unsetenv(envKey)

	value := GetEnvOrDefault(envKey, "")
	if value != envValue {
		t.Errorf("unexpected result = %v, want %v", value, envValue)
	}
}

func TestGetEnvDefault(t *testing.T) {
	envKey := "TEST_ENV_KEY_DEFAULT"
	defaultEnvValue := "TEST_ENV_VALUE_DEFAULT"

	value := GetEnvOrDefault(envKey, defaultEnvValue)
	if value != defaultEnvValue {
		t.Errorf("unexpected result = %v, want %v", value, defaultEnvValue)
	}
}
