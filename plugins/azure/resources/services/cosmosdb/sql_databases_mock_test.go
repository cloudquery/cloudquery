package cosmosdb

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

var fakeResourceGroup = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/cqprovidertest"

func buildCosmosDBSQLDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	accountSvc := mocks.NewMockCosmosDBAccountClient(ctrl)
	sqlSvc := mocks.NewMockCosmosDBSQLClient(ctrl)

	s := services.Services{
		CosmosDb: services.CosmosDbClient{
			Accounts: accountSvc,
			SQL:      sqlSvc,
		},
	}
	account := documentdb.DatabaseAccountGetResults{}
	err := faker.FakeData(&account)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	account.ID = &fakeResourceGroup

	accountSvc.EXPECT().List(gomock.Any()).Return(
		documentdb.DatabaseAccountsListResult{Value: &[]documentdb.DatabaseAccountGetResults{account}}, nil,
	)

	resource := documentdb.SQLDatabaseGetPropertiesResource{}
	err = faker.FakeDataSkipFields(&resource, []string{"Ts"})
	resource.Ts = float64(1)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	options := documentdb.SQLDatabaseGetPropertiesOptions{}
	err = faker.FakeData(&options)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	sqlDatabase := documentdb.SQLDatabaseGetResults{}
	err = faker.FakeDataSkipFields(&sqlDatabase, []string{"SQLDatabaseGetProperties"})
	sqlDatabase.SQLDatabaseGetProperties = &documentdb.SQLDatabaseGetProperties{
		Options:  &options,
		Resource: &resource,
	}
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	sqlSvc.EXPECT().ListSQLDatabases(gomock.Any(), "test", *account.Name).Return(
		documentdb.SQLDatabaseListResult{Value: &[]documentdb.SQLDatabaseGetResults{sqlDatabase}}, nil,
	)

	return s
}

func TestCosmosDBSQLDatabases(t *testing.T) {
	client.AzureMockTestHelper(t, CosmosDBSqlDatabases(), buildCosmosDBSQLDatabasesMock, client.TestOptions{})
}
