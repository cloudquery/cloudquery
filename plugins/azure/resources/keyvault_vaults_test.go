package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault71 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

var fakeResourceGroup = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/cqprovidertest"

func buildKeyVaultMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	v := mocks.NewMockVaultClient(ctrl)
	k71 := mocks.NewMockKeyVault71Client(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			KeyVault71: k71,
			Vaults:     v,
		},
	}
	vault := keyvault.Vault{}
	err := faker.FakeData(&vault)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	vaultName := fakeResourceGroup
	vault.ID = &vaultName
	vaultPage := keyvault.NewVaultListResultPage(keyvault.VaultListResult{Value: &[]keyvault.Vault{vault}}, func(ctx context.Context, result keyvault.VaultListResult) (keyvault.VaultListResult, error) {
		return keyvault.VaultListResult{}, nil
	})
	v.EXPECT().ListBySubscription(gomock.Any(), gomock.Any()).Return(vaultPage, nil)

	key := keyvault71.KeyItem{}
	if err := faker.FakeData(&key); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	keyPage := keyvault71.NewKeyListResultPage(keyvault71.KeyListResult{Value: &[]keyvault71.KeyItem{key}}, func(ctx context.Context, result keyvault71.KeyListResult) (keyvault71.KeyListResult, error) {
		return keyvault71.KeyListResult{}, nil
	})
	k71.EXPECT().GetKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(keyPage, nil)

	var secret keyvault71.SecretItem
	if err := faker.FakeData(&secret); err != nil {
		t.Fatal(err)
	}
	k71.EXPECT().GetSecrets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		keyvault71.NewSecretListResultPage(
			keyvault71.SecretListResult{Value: &[]keyvault71.SecretItem{secret}},
			func(context.Context, keyvault71.SecretListResult) (keyvault71.SecretListResult, error) {
				return keyvault71.SecretListResult{}, nil
			},
		), nil,
	)

	return s
}

func TestKeyVaultVaults(t *testing.T) {
	azureTestHelper(t, resources.KeyvaultVaults(), buildKeyVaultMock)
}
