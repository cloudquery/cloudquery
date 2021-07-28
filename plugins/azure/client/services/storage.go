package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

type StorageClient struct {
	Accounts     StorageAccountClient
	BlobServices StorageBlobServicesClient
	Containers   StorageContainerClient
}

func NewStorageClient(subscriptionId string, auth autorest.Authorizer) StorageClient {
	accounts := storage.NewAccountsClient(subscriptionId)
	accounts.Authorizer = auth
	blobServices := storage.NewBlobServicesClient(subscriptionId)
	blobServices.Authorizer = auth
	containers := storage.NewBlobContainersClient(subscriptionId)
	containers.Authorizer = auth
	return StorageClient{
		Accounts:     accounts,
		BlobServices: blobServices,
		Containers:   containers,
	}
}

type StorageAccountClient interface {
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
}

type StorageBlobServicesClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceItems, err error)
}

type StorageContainerClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, include storage.ListContainersInclude) (result storage.ListContainerItemsPage, err error)
}
