// Auto generated code - DO NOT EDIT.

package storage

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func createContainersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStorageContainersClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Containers: mockClient,
		},
	}

	data := storage.ListContainerItem{}
	require.Nil(t, faker.FakeObject(&data))

	result := storage.NewListContainerItemsPage(storage.ListContainerItems{Value: &[]storage.ListContainerItem{data}}, func(ctx context.Context, result storage.ListContainerItems) (storage.ListContainerItems, error) {
		return storage.ListContainerItems{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test", "test", "", "", gomock.Any()).Return(result, nil)
	return s
}
