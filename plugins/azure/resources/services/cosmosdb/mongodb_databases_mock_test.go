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

func buildCosmosDBMongoDBDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	accountSvc := mocks.NewMockCosmosDBAccountClient(ctrl)
	mongoDBSvc := mocks.NewMockCosmosDBMongoDBClient(ctrl)

	s := services.Services{
		CosmosDb: services.CosmosDbClient{
			Accounts: accountSvc,
			MongoDB:  mongoDBSvc,
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

	resource := documentdb.MongoDBDatabaseGetPropertiesResource{}
	err = faker.FakeDataSkipFields(&resource, []string{"Ts"})
	resource.Ts = float64(1)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	options := documentdb.MongoDBDatabaseGetPropertiesOptions{}
	err = faker.FakeData(&options)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	mongoDBDatabase := documentdb.MongoDBDatabaseGetResults{}
	err = faker.FakeDataSkipFields(&mongoDBDatabase, []string{"MongoDBDatabaseGetProperties"})
	mongoDBDatabase.MongoDBDatabaseGetProperties = &documentdb.MongoDBDatabaseGetProperties{
		Options:  &options,
		Resource: &resource,
	}
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	mongoDBSvc.EXPECT().ListMongoDBDatabases(gomock.Any(), "test", *account.Name).Return(
		documentdb.MongoDBDatabaseListResult{Value: &[]documentdb.MongoDBDatabaseGetResults{mongoDBDatabase}}, nil,
	)

	return s
}

func TestCosmosDBMongoDBDatabases(t *testing.T) {
	client.AzureMockTestHelper(t, CosmosDBMongoDBDatabases(), buildCosmosDBMongoDBDatabasesMock, client.TestOptions{})
}
