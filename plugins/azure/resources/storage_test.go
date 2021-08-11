package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues"
)

func buildStorageMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	acc := mocks.NewMockStorageAccountClient(ctrl)
	cont := mocks.NewMockStorageContainerClient(ctrl)
	blob := mocks.NewMockStorageBlobServicesClient(ctrl)
	blobProps := mocks.NewMockStorageBlobServicePropertiesClient(ctrl)
	queueProps := mocks.NewMockStorageQueueServicePropertiesClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Accounts:     acc,
			BlobServices: blob,
			Containers:   cont,
			NewBlobServiceProperties: func(autorest.Authorizer) services.StorageBlobServicePropertiesClient {
				return blobProps
			},
			NewQueueServiceProperties: func(autorest.Authorizer) services.StorageQueueServicePropertiesClient {
				return queueProps
			},
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

	// expect accounts ListKeys to be called several times
	keyValue := "dGVzdGtleQ=="
	accountKey := storage.AccountKey{
		KeyName: new(string),
		Value:   &keyValue,
	}
	acc.EXPECT().ListKeys(gomock.Any(), "test", *account.Name, storage.ListKeyExpand("")).Return(
		storage.AccountListKeysResult{Keys: &[]storage.AccountKey{accountKey}}, nil,
	).Times(2)

	// expect a call to blob GetServiceProperties
	var propsLogging accounts.GetServicePropertiesResult
	if err := faker.FakeData(&propsLogging); err != nil {
		t.Fatal(err)
	}
	blobProps.EXPECT().GetServiceProperties(gomock.Any(), *account.Name).Return(propsLogging, nil)

	// expect a call to queue GetServiceProperties
	var queueServiceProps queues.StorageServicePropertiesResponse
	if err := faker.FakeData(&queueServiceProps); err != nil {
		t.Fatal(err)
	}
	queueProps.EXPECT().GetServiceProperties(gomock.Any(), *account.Name).Return(queueServiceProps, nil)

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
