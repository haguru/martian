package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/pelletier/go-toml"
)

type AppConfig struct {
	ServiceName string
	Config      ServiceConfig
}

type ServiceConfig struct {
	Vault       VaultConfig
	RedisConfig RedisConfiguration
}

type VaultConfig struct {
	TokenFile   string
	Host        string
	Port        int
	HttpTimeout string
}
type RedisConfiguration struct {
	Host      string
	Port      int
	BatchSize int
	Timeout   time.Duration
	Password  string //consider using vault of osme type of secure way to get it maybe create an interface for this ATM
}

type MongoConfiguration struct {
	Host      string
	Port      int
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

	if reflect.DeepEqual(appConfig.Config.RedisConfig, RedisConfiguration{}) {
		return errors.New("redis configuration hase not been set")
	}
	return nil
}
