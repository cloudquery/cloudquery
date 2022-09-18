// Auto generated code - DO NOT EDIT.

package cosmosdb

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
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
			Accounts:         mockClient,
			MongoDBDatabases: createMongoDBDatabasesMock(t, ctrl).CosmosDB.MongoDBDatabases,
			SQLDatabases:     createSQLDatabasesMock(t, ctrl).CosmosDB.SQLDatabases,
		},
	}

	data := documentdb.DatabaseAccountGetResults{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := documentdb.DatabaseAccountsListResult{Value: &[]documentdb.DatabaseAccountGetResults{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
