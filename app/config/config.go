package config

import (
	"fmt"
	"os"

	"github.com/jo-hoe/ai-assistent/app/aiclient"

	"gopkg.in/yaml.v3"
)

type ConfigDescription struct {
	AIClientConfigDescription []AIClientConfigDescription `yaml:"aiclients"`
}

type AIClientConfigDescription struct {
	Type       string            `yaml:"type"`
	Properties map[string]string `yaml:",inline"`
}

type Config struct {
	AIClients []aiclient.AIClient
}

func NewConfig(path string) (config *Config, err error) {
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
	}, nil
}
