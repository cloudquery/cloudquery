// Auto generated code - DO NOT EDIT.

package cosmosdb

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
)

func createMongoDBDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCosmosDBMongoDBDatabasesClient(ctrl)
	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			MongoDBDatabases: mockClient,
		},
	}

	data := documentdb.MongoDBDatabaseGetResults{}
	require.Nil(t, faker.FakeObject(&data))

	result := documentdb.MongoDBDatabaseListResult{Value: &[]documentdb.MongoDBDatabaseGetResults{data}}

	mockClient.EXPECT().ListMongoDBDatabases(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
