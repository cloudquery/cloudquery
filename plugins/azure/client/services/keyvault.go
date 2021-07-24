package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault71 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/Azure/go-autorest/autorest"
	auth2 "github.com/Azure/go-autorest/autorest/azure/auth"
)

type KeyVaultClient struct {
	KeyVault71 KeyVault71Client
	Vaults     VaultClient
}

func NewKeyVaultClient(subscriptionId string, auth autorest.Authorizer) KeyVaultClient {
	kv71 := keyvault71.New()
	a, _ := auth2.NewAuthorizerFromEnvironmentWithResource("https://vault.azure.net")
	kv71.Authorizer = a
	vaultSvc := keyvault.NewVaultsClient(subscriptionId)
	vaultSvc.Authorizer = auth
	return KeyVaultClient{
		Vaults:     vaultSvc,
		KeyVault71: kv71,
	}
}

type VaultClient interface {
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
}

type KeyVault71Client interface {
	GetKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.KeyListResultPage, err error)
	GetSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.SecretListResultPage, err error)
}
