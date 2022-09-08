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
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := account.DataLakeStoreAccountListResult{Value: &[]account.DataLakeStoreAccountBasic{data}}

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).Return(result, nil)
	return s
}
