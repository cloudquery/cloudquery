// Auto generated code - DO NOT EDIT.

package keyvault

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
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
		},
	}

	data := keyvault.Vault{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := keyvault.NewVaultListResultPage(keyvault.VaultListResult{Value: &[]keyvault.Vault{data}}, func(ctx context.Context, result keyvault.VaultListResult) (keyvault.VaultListResult, error) {
		return keyvault.VaultListResult{}, nil
	})

	maxResults := int32(1000)
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &maxResults).Return(result, nil)
	return s
}
