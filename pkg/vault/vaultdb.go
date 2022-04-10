package vault

type VaultDB struct {
	SecretsPath string
	client      VaultClientInterface
}

func NewVaultDB(client VaultClientInterface, path string) VaultInterface {
	vaultClient := &VaultDB{
		client:      client,
		SecretsPath: path,
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

func (vdb *VaultDB) SetSecrets(secretid string, secret map[string]interface{}) error {
	wrappedSecret := map[string]interface{}{"data": secret}
	_, err := vdb.client.Set(vdb.SecretsPath+secretid, wrappedSecret)
	return err
}

func (vdb *VaultDB) DeleteSecrets(secretid string) error {
	_, err := vdb.client.Delete(vdb.SecretsPath + secretid)
	return err
}
