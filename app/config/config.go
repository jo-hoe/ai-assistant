package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/jo-hoe/ai-assistent/app/aiclient"
	"github.com/jo-hoe/ai-assistent/app/common"

	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_CONFIG_PATH_KEY = "CONFIG_PATH"
	DEFAULT_CONFIG_PATH     = "config.yaml"
)

type ConfigDescription struct {
	AIClientConfigDescription []AIClientConfigDescription `yaml:"aiclients"`
	Port                      int                         `yaml:"port"`
}

type AIClientConfigDescription struct {
	Type       string            `yaml:"type"`
	Properties map[string]string `yaml:",inline"`
}

type Config struct {
	Port      int
	AIClients aiclient.AIClients
}

var (
	once sync.Once
	// singleton of config
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		var err error
		configInstance, err = loadConfig()
		if err != nil {
			panic(err)
		}
	})

	return configInstance
}

func loadConfig() (config *Config, err error) {
	configPath := common.GetEnvOrDefault(DEFAULT_CONFIG_PATH_KEY, DEFAULT_CONFIG_PATH)
	return newConfig(configPath)
}

func newConfig(path string) (config *Config, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configDescription ConfigDescription
	err = yaml.Unmarshal(data, &configDescription)
	if err != nil {
		return nil, err
	}

	var clients []aiclient.AIClient
	for _, clientsConfig := range configDescription.AIClientConfigDescription {
		var client aiclient.AIClient

		switch clientsConfig.Type {
		case aiclient.MOCK_CLIENT_TYPE_NAME:
			client, err = aiclient.NewMockClientFromMap(clientsConfig.Properties)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown service type: %s", clientsConfig.Type)
		}

		clients = append(clients, client)
	}

	return &Config{
		AIClients: clients,
		Port:      configDescription.Port,
	}, nil
}
