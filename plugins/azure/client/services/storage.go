package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

type StorageClient struct {
	Containers StorageContainerClient
	Accounts   StorageAccountClient
}

func NewStorageClient(subscriptionId string, auth autorest.Authorizer) StorageClient {
	containers := storage.NewBlobContainersClient(subscriptionId)
	containers.Authorizer = auth
	accounts := storage.NewAccountsClient(subscriptionId)
	accounts.Authorizer = auth
	return StorageClient{
		Containers: containers,
		Accounts:   accounts,
	}
}

type StorageContainerClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, include storage.ListContainersInclude) (result storage.ListContainerItemsPage, err error)
}

type StorageAccountClient interface {
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
}
