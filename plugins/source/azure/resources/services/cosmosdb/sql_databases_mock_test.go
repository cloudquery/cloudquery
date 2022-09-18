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

func createSQLDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCosmosDBSQLDatabasesClient(ctrl)
	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			SQLDatabases: mockClient,
		},
	}

	data := documentdb.SQLDatabaseGetResults{}
	require.Nil(t, faker.FakeObject(&data))

	result := documentdb.SQLDatabaseListResult{Value: &[]documentdb.SQLDatabaseGetResults{data}}

	mockClient.EXPECT().ListSQLDatabases(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
