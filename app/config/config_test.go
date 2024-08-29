package config

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jo-hoe/ai-assistent/app/aiclient"
)

const envFileDir = "testdata"
const envFileName = "testconfig.yaml"

func TestNewConfig(t *testing.T) {
	configPath := getTestConfigPath(t)
	os.Setenv(DEFAULT_CONFIG_PATH_KEY, configPath)
	defer os.Unsetenv(DEFAULT_CONFIG_PATH_KEY)

	config := GetConfig()

	if config == nil {
		t.Error("config is nil")
		return
	}

	expectedPort := 8080
	if config.Port != expectedPort {
		t.Errorf("unexpected result = %v, want %v", config.Port, expectedPort)
	}

	if config.AIClients == nil {
		t.Error("config.AIClients is nil")
	}

	expectedMockConfig := aiclient.NewMockClient([]string{"42", "this is an answer"}, 500, "error")
	if !reflect.DeepEqual(config.AIClients[0], expectedMockConfig) {
		t.Errorf("unexpected result = %v, want %v", config.AIClients[0], expectedMockConfig)
	}
}

func getTestConfigPath(t *testing.T) string {
	current_directory, error := os.Getwd()
	if error != nil {
		t.Error(error)
	}

	appFolderPath := filepath.Dir(current_directory)
	workingDirectoryPath := filepath.Dir(appFolderPath)
	return path.Join(workingDirectoryPath, envFileDir, envFileName)
}
