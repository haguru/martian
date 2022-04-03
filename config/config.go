package config

import (
	"errors"
	"reflect"
	"time"

	"github.com/haguru/martian/pkg/db"
	"github.com/hashicorp/vault/api"
)

type App struct {
	config ServiceConfig
}


type ServiceConfig struct {
	Vault VaultConfig
	RedisConfig db.Configuration
}


type VaultConfig struct {
	ApiConfig  api.Config
	HttpTimeout time.Duration
}

// double check this linting: avoid using reflect.DeepEqual with errorsdeepequalerrors
func (appConfig *App) Validate() error {
	if reflect.DeepEqual(appConfig.config.Vault, api.Config{}) {
		return errors.New("vault configuration hase not been set")
	}

	if reflect.DeepEqual(appConfig.config.RedisConfig, db.Configuration{}) {
		return errors.New("redis configuration hase not been set")
	}
	return nil
}
