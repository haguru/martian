package db

import (
	"fmt"
	"net/http"
	"time"

	"github.com/haguru/martian/config"
	"github.com/hashicorp/vault/api"
)

type VaultClient struct {
	client *api.Client
}


// TODO: set token needs to happen. also change config to hous all info needed
func NewVaultClient(vconfig config.VaultConfig) (VaultClientInterface,error) {
	vconfig.ApiConfig.HttpClient = &http.Client{
		Timeout: vconfig.HttpTimeout * time.Second,
	}
	client, err := api.NewClient(&vconfig.ApiConfig)
	if err != nil{
		return nil, fmt.Errorf("unable to create vault client: %v",err)
	}

	return &VaultClient{client: client},nil
}

func (vc *VaultClient) Set(path string, data map[string]interface{}) (*api.Secret, error){
	secret, err := vc.client.Logical().Write(path, data)
	if err != nil{
		return nil,fmt.Errorf("unable to set secret in vailt: %v",err)
	}
	return secret, nil
}

func (vc *VaultClient) Get(path string) (*api.Secret, error){
	secret, err := vc.client.Logical().Read(path)
	if err != nil{
		return nil,fmt.Errorf("unable to get secret from vailt: %v",err)
	}
	return secret,nil
}

func (vc *VaultClient) Delete(path string) (*api.Secret, error){
	secret, err := vc.client.Logical().Delete(path)
	if err != nil{
		return nil,fmt.Errorf("unable to delete secret in vailt: %v",err)
	}
	return secret, nil
}
