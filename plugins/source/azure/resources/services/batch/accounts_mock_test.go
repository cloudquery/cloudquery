// Auto generated code - DO NOT EDIT.

package batch

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
)

func TestBatchAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccountsMock)
}

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockBatchAccountsClient(ctrl)
	s := services.Services{
		Batch: services.BatchClient{
			Accounts: mockClient,
		},
	}

	data := batch.Account{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := batch.NewAccountListResultPage(batch.AccountListResult{Value: &[]batch.Account{data}}, func(ctx context.Context, result batch.AccountListResult) (batch.AccountListResult, error) {
		return batch.AccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
