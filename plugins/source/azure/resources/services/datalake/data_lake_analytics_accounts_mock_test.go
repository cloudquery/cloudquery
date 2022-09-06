// Auto generated code - DO NOT EDIT.

package datalake

import (
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

func TestDataLakeDataLakeAnalyticsAccounts(t *testing.T) {
	client.AzureMockTestHelper(t, DataLakeAnalyticsAccounts(), createDataLakeAnalyticsAccountsMock, client.TestOptions{})
}

func createDataLakeAnalyticsAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockDataLakeDataLakeAnalyticsAccountsClient(ctrl)
	s := services.Services{
		DataLake: services.DataLakeClient{
			DataLakeAnalyticsAccounts: mockClient,
		},
	}

	data := account.DataLakeAnalyticsAccountBasic{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := account.DataLakeAnalyticsAccountListResult{Value: &[]account.DataLakeAnalyticsAccountBasic{data}}

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).Return(result, nil)
	return s
}
