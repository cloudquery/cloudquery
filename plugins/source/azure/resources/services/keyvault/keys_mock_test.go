// Auto generated code - DO NOT EDIT.

package keyvault

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
)

func createKeysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultKeysClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Keys: mockClient,
		},
	}

	data := keyvault.KeyItem{}
	require.Nil(t, faker.FakeObject(&data))

	result := keyvault.NewKeyListResultPage(keyvault.KeyListResult{Value: &[]keyvault.KeyItem{data}}, func(ctx context.Context, result keyvault.KeyListResult) (keyvault.KeyListResult, error) {
		return keyvault.KeyListResult{}, nil
	})

	maxResults := int32(25)
	mockClient.EXPECT().GetKeys(gomock.Any(), "test", &maxResults).Return(result, nil)
	return s
}
