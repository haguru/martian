package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/haguru/martian/pkg/redis"
	"github.com/hashicorp/vault/api"
	"github.com/pelletier/go-toml"
)

type AppConfig struct {
	ServiceName string
	Config ServiceConfig
}

type ServiceConfig struct {
	Vault       VaultConfig
	RedisConfig redis.Configuration
}

type VaultConfig struct {
	TokenFile   string
	Host        string
	Port        int
	HttpTimeout string
}

func LoadConfig() (*AppConfig, error) {
	config := &AppConfig{}
	file, err := os.ReadFile("./res/configuration.toml")

	if err != nil {
		return nil, fmt.Errorf("unable to read in configuration toml: %v", err)
	}
	err = toml.Unmarshal(file, config)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal configuration toml: %v", err)
	}
	return config, nil
}

// double check this linting: avoid using reflect.DeepEqual with errorsdeepequalerrors
func Validate(appConfig *AppConfig) error {
	if reflect.DeepEqual(appConfig.Config.Vault, api.Config{}) {
		return errors.New("vault configuration hase not been set")
	}

	if reflect.DeepEqual(appConfig.Config.RedisConfig, redis.Configuration{}) {
		return errors.New("redis configuration hase not been set")
	}
	return nil
}
