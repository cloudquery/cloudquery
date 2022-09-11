// Auto generated code - DO NOT EDIT.

package datalake

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
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	getData := account.DataLakeAnalyticsAccount{}
	require.Nil(t, faker.FakeData(&getData, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := account.NewDataLakeAnalyticsAccountListResultPage(account.DataLakeAnalyticsAccountListResult{Value: &[]account.DataLakeAnalyticsAccountBasic{data}}, func(ctx context.Context, result account.DataLakeAnalyticsAccountListResult) (account.DataLakeAnalyticsAccountListResult, error) {
		return account.DataLakeAnalyticsAccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).Return(result, nil)

	mockClient.EXPECT().Get(gomock.Any(), "test", *data.Name).Return(getData, nil)
	return s
}
