// Auto generated code - DO NOT EDIT.

package batch

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
)

func TestBatchAccounts(t *testing.T) {
	client.AzureMockTestHelper(t, Accounts(), createAccountsMock, client.TestOptions{})
}

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockBatchAccountsClient(ctrl)
	s := services.Services{
		Batch: services.BatchClient{
			Accounts: mockClient,
		},
	}

	data := batch.Account{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := batch.NewAccountListResultPage(batch.AccountListResult{Value: &[]batch.Account{data}}, func(ctx context.Context, result batch.AccountListResult) (batch.AccountListResult, error) {
		return batch.AccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
