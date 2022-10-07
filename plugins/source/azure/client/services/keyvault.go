//go:generate mockgen -destination=./mocks/keyvault.go -package=mocks . KeyVaultVaultsClient,KeyVaultManagedHsmsClient,KeyVaultKeysClient,KeyVaultSecretsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault71 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	hsm "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	"github.com/Azure/go-autorest/autorest"
	auth2 "github.com/Azure/go-autorest/autorest/azure/auth"
)

type KeyVaultClient struct {
	Keys        KeyVaultKeysClient
	Vaults      KeyVaultVaultsClient
	ManagedHsms KeyVaultManagedHsmsClient
	Secrets     KeyVaultSecretsClient
}

type KeyVaultVaultsClient interface {
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
}

type KeyVaultKeysClient interface {
	GetKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.KeyListResultPage, err error)
}

type KeyVaultSecretsClient interface {
	GetSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.SecretListResultPage, err error)
}

type KeyVaultManagedHsmsClient interface {
	ListBySubscription(ctx context.Context, top *int32) (result hsm.ManagedHsmListResultPage, err error)
}

func NewKeyVaultClient(subscriptionId string, auth autorest.Authorizer) (KeyVaultClient, error) {
	kv71 := keyvault71.New()
	// The audience for keyvault71 should be different so we need to request a new token
	// https://stackoverflow.com/questions/60216664/update-azure-keyvault-secret-through-azure-api
	a, err := auth2.NewAuthorizerFromCLIWithResource("https://vault.azure.net")
	if err != nil {
		a, err = auth2.NewAuthorizerFromEnvironmentWithResource("https://vault.azure.net")
		if err != nil {
			return KeyVaultClient{}, err
		}
	}
	kv71.Authorizer = a

	vaultSvc := keyvault.NewVaultsClient(subscriptionId)
	vaultSvc.Authorizer = auth

	vhsm := hsm.NewManagedHsmsClient(subscriptionId)
	vhsm.Authorizer = auth

	return KeyVaultClient{
		Vaults:      vaultSvc,
		Keys:        kv71,
		Secrets:     kv71,
		ManagedHsms: vhsm,
	}, nil
}
