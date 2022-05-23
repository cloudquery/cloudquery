package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues"
)

type StorageClient struct {
	Accounts                  StorageAccountClient
	BlobServices              StorageBlobServicesClient
	Containers                StorageContainerClient
	NewBlobServiceProperties  func(autorest.Authorizer) StorageBlobServicePropertiesClient
	NewQueueServiceProperties func(autorest.Authorizer) StorageQueueServicePropertiesClient
}

type StorageAccountClient interface {
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string, expand storage.ListKeyExpand) (result storage.AccountListKeysResult, err error)
}

type StorageBlobServicesClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceItems, err error)
}

type StorageContainerClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, include storage.ListContainersInclude) (result storage.ListContainerItemsPage, err error)
}

type StorageBlobServicePropertiesClient interface {
	GetServiceProperties(ctx context.Context, accountName string) (result accounts.GetServicePropertiesResult, err error)
}

type StorageQueueServicePropertiesClient interface {
	GetServiceProperties(ctx context.Context, accountName string) (result queues.StorageServicePropertiesResponse, err error)
}

func NewStorageClient(subscriptionId string, auth autorest.Authorizer) StorageClient {
	accs := storage.NewAccountsClient(subscriptionId)
	accs.Authorizer = auth
	blobServices := storage.NewBlobServicesClient(subscriptionId)
	blobServices.Authorizer = auth
	containers := storage.NewBlobContainersClient(subscriptionId)
	containers.Authorizer = auth
	return StorageClient{
		Accounts:     accs,
		BlobServices: blobServices,
		Containers:   containers,
		NewBlobServiceProperties: func(auth autorest.Authorizer) StorageBlobServicePropertiesClient {
			client := accounts.New()
			client.Authorizer = auth
			return client
		},
		NewQueueServiceProperties: func(auth autorest.Authorizer) StorageQueueServicePropertiesClient {
			client := queues.New()
			client.Authorizer = auth
			return client
		},
	}
}
