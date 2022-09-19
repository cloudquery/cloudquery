// Auto generated code - DO NOT EDIT.

package keyvault

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
)

func TestKeyVaultVaults(t *testing.T) {
	client.MockTestHelper(t, Vaults(), createVaultsMock)
}

func createVaultsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultVaultsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Vaults: mockClient,
			Keys:   createKeysMock(t, ctrl).KeyVault.Keys,
		},
	}

	data := keyvault.Vault{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := keyvault.NewVaultListResultPage(keyvault.VaultListResult{Value: &[]keyvault.Vault{data}}, func(ctx context.Context, result keyvault.VaultListResult) (keyvault.VaultListResult, error) {
		return keyvault.VaultListResult{}, nil
	})

	maxResults := int32(1000)
	vaultURI := "test"
	data.Properties.VaultURI = &vaultURI
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &maxResults).Return(result, nil)
	return s
}
