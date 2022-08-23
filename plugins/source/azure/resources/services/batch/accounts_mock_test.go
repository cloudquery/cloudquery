package batch

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildBatchAccountsClientMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockAccountsClient(ctrl)
	var acc batch.Account
	if err := faker.FakeData(&acc); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any()).Return(
		batch.NewAccountListResultPage(
			batch.AccountListResult{Value: &[]batch.Account{acc}},
			func(c context.Context, lr batch.AccountListResult) (batch.AccountListResult, error) {
				return batch.AccountListResult{}, nil
			},
		),
		nil,
	)
	return services.Services{
		Batch: services.BatchClient{
			Accounts: m,
		},
	}
}

func TestBatchAccounts(t *testing.T) {
	client.AzureMockTestHelper(t, BatchAccounts(), buildBatchAccountsClientMock, client.TestOptions{})
}
