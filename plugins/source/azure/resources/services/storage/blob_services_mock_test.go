// Auto generated code - DO NOT EDIT.

package storage

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func createBlobServicesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStorageBlobServicesClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			BlobServices: mockClient,
		},
	}

	data := storage.BlobServiceProperties{}
	require.Nil(t, faker.FakeObject(&data))

	result := storage.BlobServiceItems{Value: &[]storage.BlobServiceProperties{data}}

	mockClient.EXPECT().List(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
