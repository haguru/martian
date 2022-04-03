package config

import (
	"errors"
	"reflect"

	"github.com/haguru/martian/pkg/db"
	"github.com/hashicorp/vault/api"
)

type App struct {
	config ServiceConfig
}

type ServiceConfig struct {
	ValutConfig api.Config
	RedisConfig db.Configuration
}

// double check this linting: avoid using reflect.DeepEqual with errorsdeepequalerrors
func (appConfig *App) Validate() error {
	if reflect.DeepEqual(appConfig.config.ValutConfig, api.Config{}) {
		return errors.New("vault configuration hase not been set")
	}

	if reflect.DeepEqual(appConfig.config.RedisConfig, db.Configuration{}) {
		return errors.New("redis configuration hase not been set")
	}
	return nil
}
