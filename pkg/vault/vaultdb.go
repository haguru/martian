package db

type VaultDB struct {
	SecretsPath string
	client      VaultClientInterface
}

func NewVaultDB(client VaultClientInterface) VaultInterface {
	vaultClient := &VaultDB{
		client: client,
	}

	return vaultClient
}

func (vdb *VaultDB) GetSecrets(secretid string) (map[string]interface{}, error) {
	secret, err := vdb.client.Get(vdb.SecretsPath + secretid)
	if err != nil {
		return nil, err
	}

	return secret.Data, nil
}

func (vdb *VaultDB) SetSecrets(secretName string, data map[string]interface{}) error{
	_,err := vdb.client.Set(vdb.SecretsPath+secretName, data)
	return err
}

func (vdb *VaultDB) DeleteSecrets(secretid string) error{
	_,err := vdb.client.Delete(vdb.SecretsPath + secretid)
	return err
}