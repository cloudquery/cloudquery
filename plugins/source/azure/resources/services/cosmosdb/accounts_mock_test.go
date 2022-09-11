// Auto generated code - DO NOT EDIT.

package cosmosdb

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
)

func TestCosmosDBAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccountsMock)
}

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCosmosDBAccountsClient(ctrl)
	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			Accounts: mockClient,
		},
	}

	data := documentdb.DatabaseAccountGetResults{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := documentdb.DatabaseAccountsListResult{Value: &[]documentdb.DatabaseAccountGetResults{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
