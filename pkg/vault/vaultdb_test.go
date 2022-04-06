package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/haguru/martian/mocks"
	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/mock"
)

func TestVaultDB_GetSecrets(t *testing.T) {
	type fields struct {
		SecretsPath string
		client      *mocks.VaultClientInterface
	}
	type args struct {
		secretid string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		ClientGetErr    error
		ClientGetReturn *api.Secret
		want            map[string]interface{}
		wantErr         bool
	}{
		{
			name: "Successful Get",
			fields: fields{
				client:      &mocks.VaultClientInterface{},
				SecretsPath: "path/to/Secret",
			},
			args: args{
				secretid: "testID",
			},
			want:         map[string]interface{}{"testID": "secret"},
			wantErr:      false,
			ClientGetErr: nil,
			ClientGetReturn: &api.Secret{
				Data: map[string]interface{}{"testID": "secret"},
			},
		},
		{
			name: "Client error",
			fields: fields{
				client:      &mocks.VaultClientInterface{},
				SecretsPath: "path/to/Secret",
			},
			args: args{
				secretid: "testID",
			},
			want:            nil,
			wantErr:         true,
			ClientGetErr:    fmt.Errorf("fail"),
			ClientGetReturn: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.client.On("Get", mock.Anything, mock.Anything).Return(tt.ClientGetReturn, tt.ClientGetErr)
			vdb := &VaultDB{
				SecretsPath: tt.fields.SecretsPath,
				client:      tt.fields.client,
			}
			got, err := vdb.GetSecrets(tt.args.secretid)
			if (err != nil) != tt.wantErr {
				t.Errorf("VaultDB.GetSecrets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultDB.GetSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDB_SetSecrets(t *testing.T) {
	type fields struct {
		SecretsPath string
		client      *mocks.VaultClientInterface
	}
	type args struct {
		secretName string
		data       map[string]interface{}
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		ClientSetReturn *api.Secret
		ClientSetErr    error
		wantErr         bool
	}{
		{
			name: "Successful set",
			fields: fields{
				client:      &mocks.VaultClientInterface{},
				SecretsPath: "path/to/Secret",
			},
			args: args{
				secretName: "testSecretName",
				data:       map[string]interface{}{"test1": "Secret"},
			},
			wantErr:         false,
			ClientSetReturn: &api.Secret{},
			ClientSetErr:    nil,
		},
		{
			name: "Client error",
			fields: fields{
				client:      &mocks.VaultClientInterface{},
				SecretsPath: "path/to/Secret",
			},
			args: args{
				secretName: "testSecretName",
				data:       map[string]interface{}{"test1": "Secret"},
			},
			wantErr:         true,
			ClientSetReturn: &api.Secret{},
			ClientSetErr:    fmt.Errorf("fail"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.client.On("Set", mock.Anything, mock.Anything).Return(tt.ClientSetReturn, tt.ClientSetErr)
			vdb := &VaultDB{
				SecretsPath: tt.fields.SecretsPath,
				client:      tt.fields.client,
			}
			if err := vdb.SetSecrets(tt.args.secretName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("VaultDB.SetSecrets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVaultDB_DeleteSecrets(t *testing.T) {
	type fields struct {
		SecretsPath string
		client      *mocks.VaultClientInterface
	}
	type args struct {
		secretid string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		ClientDeleteErr    error
		ClientDeleteReturn *api.Secret
		wantErr            bool
	}{
		{
			name: "Successful Delete",
			fields: fields{
				client: &mocks.VaultClientInterface{},
			},
			args: args{
				secretid: "testSecretID",
			},
			wantErr:            false,
			ClientDeleteErr:    nil,
			ClientDeleteReturn: &api.Secret{},
		},
		{
			name: "Client Error",
			fields: fields{
				client: &mocks.VaultClientInterface{},
			},
			args: args{
				secretid: "testSecretID",
			},
			wantErr:            true,
			ClientDeleteErr:    fmt.Errorf("fail"),
			ClientDeleteReturn: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.client.On("Delete", mock.Anything, mock.Anything).Return(tt.ClientDeleteReturn, tt.ClientDeleteErr)
			vdb := &VaultDB{
				SecretsPath: tt.fields.SecretsPath,
				client:      tt.fields.client,
			}
			if err := vdb.DeleteSecrets(tt.args.secretid); (err != nil) != tt.wantErr {
				t.Errorf("VaultDB.DeleteSecrets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
