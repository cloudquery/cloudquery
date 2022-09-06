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

	"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
)

func TestKeyVaultManagedHSMs(t *testing.T) {
	client.AzureMockTestHelper(t, ManagedHSMs(), createManagedHSMsMock, client.TestOptions{})
}

func createManagedHSMsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultManagedHSMsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			ManagedHSMs: mockClient,
		},
	}

	data := keyvault.ManagedHsm{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := keyvault.NewManagedHsmListResultPage(keyvault.ManagedHsmListResult{Value: &[]keyvault.ManagedHsm{data}}, func(ctx context.Context, result keyvault.ManagedHsmListResult) (keyvault.ManagedHsmListResult, error) {
		return keyvault.ManagedHsmListResult{}, nil
	})

	maxResults := int32(100)
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &maxResults).Return(result, nil)
	return s
}
