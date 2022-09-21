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

func createSecretsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultSecretsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Secrets: mockClient,
		},
	}

	data := keyvault.SecretItem{}
	require.Nil(t, faker.FakeObject(&data))

	result := keyvault.NewSecretListResultPage(keyvault.SecretListResult{Value: &[]keyvault.SecretItem{data}}, func(ctx context.Context, result keyvault.SecretListResult) (keyvault.SecretListResult, error) {
		return keyvault.SecretListResult{}, nil
	})

	maxResults := int32(25)
	mockClient.EXPECT().GetSecrets(gomock.Any(), "test", &maxResults).Return(result, nil)
	return s
}
