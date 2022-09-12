// Auto generated code - DO NOT EDIT.

package storage

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func TestStorageAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccountsMock)
}

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStorageAccountsClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Accounts:     mockClient,
			Containers:   createContainersMock(t, ctrl).Storage.Containers,
			BlobServices: createBlobServicesMock(t, ctrl).Storage.BlobServices,
		},
	}

	data := storage.Account{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := storage.NewAccountListResultPage(storage.AccountListResult{Value: &[]storage.Account{data}}, func(ctx context.Context, result storage.AccountListResult) (storage.AccountListResult, error) {
		return storage.AccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
