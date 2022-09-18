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

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account"
)

func TestDataLakeAnalyticsAccounts(t *testing.T) {
	client.MockTestHelper(t, AnalyticsAccounts(), createAnalyticsAccountsMock)
}

func createAnalyticsAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockDataLakeAnalyticsAccountsClient(ctrl)
	s := services.Services{
		DataLake: services.DataLakeClient{
			AnalyticsAccounts: mockClient,
		},
	}

	data := account.DataLakeAnalyticsAccountBasic{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	getData := account.DataLakeAnalyticsAccount{}
	require.Nil(t, faker.FakeObject(&getData))

	result := account.NewDataLakeAnalyticsAccountListResultPage(account.DataLakeAnalyticsAccountListResult{Value: &[]account.DataLakeAnalyticsAccountBasic{data}}, func(ctx context.Context, result account.DataLakeAnalyticsAccountListResult) (account.DataLakeAnalyticsAccountListResult, error) {
		return account.DataLakeAnalyticsAccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).Return(result, nil)

	mockClient.EXPECT().Get(gomock.Any(), "test", *data.Name).Return(getData, nil)
	return s
}
