package cosmosdb

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCosmosDBAccountMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	accountSvc := mocks.NewMockCosmosDBAccountsClient(ctrl)

	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			Accounts: accountSvc,
		},
	}
	account := documentdb.DatabaseAccountGetResults{}
	err := faker.FakeData(&account)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	accountSvc.EXPECT().List(gomock.Any()).Return(
		documentdb.DatabaseAccountsListResult{Value: &[]documentdb.DatabaseAccountGetResults{account}}, nil,
	)

	return s
}

func TestCosmosDBAccount(t *testing.T) {
	client.AzureMockTestHelper(t, CosmosDBAccounts(), buildCosmosDBAccountMock, client.TestOptions{})
}
