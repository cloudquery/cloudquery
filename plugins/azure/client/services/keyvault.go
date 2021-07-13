package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/mgmt/keyvault"
	keyvault71 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

type KeyVaultClient struct {
	Keys    KeyClient
	Secrets SecretsClient
	Vaults  VaultClient
}

func NewKeyVaultClient(subscriptionId string, auth autorest.Authorizer) KeyVaultClient {
	keySvc := keyvault.NewKeysClient(subscriptionId)
	keySvc.Authorizer = auth
	secretsSvc := keyvault71.New()
	secretsSvc.Authorizer = auth
	vaultSvc := keyvault.NewVaultsClient(subscriptionId)
	vaultSvc.Authorizer = auth
	return KeyVaultClient{
		Vaults: vaultSvc,
		Keys:   keySvc,
	}
}

type VaultClient interface {
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
}

type KeyClient interface {
	List(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.KeyListResultPage, err error)
}

type SecretsClient interface {
	GetSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault71.SecretListResultPage, err error)
}
