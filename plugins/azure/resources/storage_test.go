package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildStorageMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	acc := mocks.NewMockStorageAccountClient(ctrl)
	cont := mocks.NewMockStorageContainerClient(ctrl)
	blob := mocks.NewMockStorageBlobServicesClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Accounts:     acc,
			BlobServices: blob,
			Containers:   cont,
		},
	}
	account := storage.Account{}
	err := faker.FakeData(&account)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testAccount"
	account.Name = &name
	account.ID = &fakeResourceGroup
	page := storage.NewAccountListResultPage(storage.AccountListResult{Value: &[]storage.Account{account}}, func(ctx context.Context, result storage.AccountListResult) (storage.AccountListResult, error) {
		return storage.AccountListResult{}, nil
	})
	acc.EXPECT().List(gomock.Any()).Return(page, nil)

	container := storage.ListContainerItem{}
	if err := faker.FakeData(&container); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	containerPage := storage.NewListContainerItemsPage(storage.ListContainerItems{
		Value: &[]storage.ListContainerItem{container}}, func(ctx context.Context, items storage.ListContainerItems) (storage.ListContainerItems, error) {
		return storage.ListContainerItems{}, nil
	})
	cont.EXPECT().List(gomock.Any(), "test", *account.Name, "", "", gomock.Any()).Return(containerPage, nil)

	var props storage.BlobServiceProperties
	if err := faker.FakeData(&props); err != nil {
		t.Fatal(err)
	}
	blob.EXPECT().List(gomock.Any(), "test", *account.Name).Return(
		storage.BlobServiceItems{Value: &[]storage.BlobServiceProperties{props}}, nil,
	)

	return s
}

func TestStorageAccounts(t *testing.T) {
	azureTestHelper(t, resources.StorageAccounts(), buildStorageMock)
}
