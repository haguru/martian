package db

import "github.com/hashicorp/vault/api"

type VaultClientInterface interface {
	Set(path string, data map[string]interface{}) (*api.Secret, error)
	Get(path string) (*api.Secret, error)
	Delete(path string) (*api.Secret, error)
}

type VaultInterface interface {
	GetSecrets(secretid string) (map[string]interface{}, error)
	SetSecrets(secretName string, data map[string]interface{}) error
	DeleteSecrets(secretid string) error
}
