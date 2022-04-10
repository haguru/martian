package vault

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/haguru/martian/config"
	"github.com/hashicorp/vault/api"
)

type VaultClient struct {
	client *api.Client
}

// TODO: set token needs to happen. also change config to hous all info needed
func NewVaultClient(vconfig *config.VaultConfig) (VaultClientInterface, error) {
	address := fmt.Sprintf("http://%v:%d", vconfig.Host, vconfig.Port)

	timeout, err := time.ParseDuration(vconfig.HttpTimeout)
	if err != nil {
		return nil, fmt.Errorf("unable to parse HttpTimeout to time.duration: %v", err)
	}

	vaultConfig := api.Config{
		Address: address,
		HttpClient: &http.Client{
			Timeout: timeout,
		},
	}

	client, err := api.NewClient(&vaultConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create vault client: %v", err)
	}

	token, err := os.ReadFile(vconfig.TokenFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read token file: %v", err)
	}
	client.SetToken(string(token))

	return &VaultClient{client: client}, nil
}

func (vc *VaultClient) Set(path string, data map[string]interface{}) (*api.Secret, error) {
	secret, err := vc.client.Logical().Write(path, data)
	if err != nil {
		return nil, fmt.Errorf("unable to set secret in vailt: %v", err)
	}
	return secret, nil
}

func (vc *VaultClient) Get(path string) (*api.Secret, error) {
	secret, err := vc.client.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("unable to get secret from vailt: %v", err)
	}
	return secret, nil
}

func (vc *VaultClient) Delete(path string) (*api.Secret, error) {
	secret, err := vc.client.Logical().Delete(path)
	if err != nil {
		return nil, fmt.Errorf("unable to delete secret in vailt: %v", err)
	}
	return secret, nil
}
