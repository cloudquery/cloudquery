package services

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/mgmt/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

type KeyVaultClient struct {
	Vaults VaultClient
	Keys   KeyClient
}

func NewKeyVaultClient(subscriptionId string, auth autorest.Authorizer) KeyVaultClient {
	vaultSvc := keyvault.NewVaultsClient(subscriptionId)
	vaultSvc.Authorizer = auth
	keySvc := keyvault.NewKeysClient(subscriptionId)
	keySvc.Authorizer = auth
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
