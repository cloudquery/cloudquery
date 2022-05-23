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
	KeyVault71 KeyVault71Client
	Vaults     VaultClient
	ManagedHSM KeyVaultManagedHSMClient
}

type VaultClient interface {
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
}

type KeysClient interface {
	// List lists the keys in the specified key vault.
	// Parameters:
	// resourceGroupName - the name of the resource group which contains the specified key vault.
	// vaultName - the name of the vault which contains the keys to be retrieved.
	List(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.KeyListResultPage, err error)
}

type KeyVault71Client interface {
	GetKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.KeyListResultPage, err error)
	GetSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.SecretListResultPage, err error)
}

type KeyVaultManagedHSMClient interface {
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
		Vaults:     vaultSvc,
		KeyVault71: kv71,
		ManagedHSM: vhsm,
	}, nil
}
