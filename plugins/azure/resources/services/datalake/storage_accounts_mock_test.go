package datalake

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDatalakeStorageAccounts(t *testing.T, ctrl *gomock.Controller) services.Services {
	ds := mocks.NewMockDataLakeStorageAccountsClient(ctrl)

	faker.SetIgnoreInterface(true)

	dataLakeStoreAccountBasic := account.DataLakeStoreAccountBasic{}
	if err := faker.FakeData(&dataLakeStoreAccountBasic); err != nil {
		t.Fatal(err)
	}
	id := client.FakeResourceGroup
	dataLakeStoreAccountBasic.ID = &id

	accounts := account.NewDataLakeStoreAccountListResultPage(
		account.DataLakeStoreAccountListResult{Value: &[]account.DataLakeStoreAccountBasic{dataLakeStoreAccountBasic}},
		func(ctx context.Context, result account.DataLakeStoreAccountListResult) (account.DataLakeStoreAccountListResult, error) {
			return account.DataLakeStoreAccountListResult{}, nil
		},
	)
	ds.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(accounts, nil)

	dataLakeStoreAccount := account.DataLakeStoreAccount{}
	if err := faker.FakeData(&dataLakeStoreAccount); err != nil {
		t.Fatal(err)
	}

	ip := faker.IPv4()
	(*dataLakeStoreAccount.FirewallRules)[0].EndIPAddress = &ip
	(*dataLakeStoreAccount.FirewallRules)[0].StartIPAddress = &ip
	dataLakeStoreAccount.ID = &id

	ds.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(dataLakeStoreAccount, nil)

	return services.Services{
		DataLake: services.DataLakeClient{DataLakeStorageAccounts: ds},
	}
}

func TestDatalakeStorageAccounts(t *testing.T) {
	client.AzureMockTestHelper(t, StorageAccounts(), buildDatalakeStorageAccounts, client.TestOptions{})
}
