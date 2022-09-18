// Auto generated code - DO NOT EDIT.

package datalake

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
)

func TestDataLakeStoreAccounts(t *testing.T) {
	client.MockTestHelper(t, StoreAccounts(), createStoreAccountsMock)
}

func createStoreAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockDataLakeStoreAccountsClient(ctrl)
	s := services.Services{
		DataLake: services.DataLakeClient{
			StoreAccounts: mockClient,
		},
	}

	data := account.DataLakeStoreAccountBasic{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	getData := account.DataLakeStoreAccount{}
	require.Nil(t, faker.FakeObject(&getData))

	result := account.NewDataLakeStoreAccountListResultPage(account.DataLakeStoreAccountListResult{Value: &[]account.DataLakeStoreAccountBasic{data}}, func(ctx context.Context, result account.DataLakeStoreAccountListResult) (account.DataLakeStoreAccountListResult, error) {
		return account.DataLakeStoreAccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).Return(result, nil)

	mockClient.EXPECT().Get(gomock.Any(), "test", *data.Name).Return(getData, nil)
	return s
}
