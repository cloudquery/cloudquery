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

	"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
)

func TestKeyVaultManagedHsms(t *testing.T) {
	client.MockTestHelper(t, ManagedHsms(), createManagedHsmsMock)
}

func createManagedHsmsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultManagedHsmsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			ManagedHsms: mockClient,
		},
	}

	data := keyvault.ManagedHsm{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := keyvault.NewManagedHsmListResultPage(keyvault.ManagedHsmListResult{Value: &[]keyvault.ManagedHsm{data}}, func(ctx context.Context, result keyvault.ManagedHsmListResult) (keyvault.ManagedHsmListResult, error) {
		return keyvault.ManagedHsmListResult{}, nil
	})

	maxResults := int32(100)
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &maxResults).Return(result, nil)
	return s
}
