//go:generate mockgen -destination=./mocks/storage.go -package=mocks . StorageAccountsClient,StorageBlobServicesClient,StorageContainersClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues"
)

type StorageClient struct {
	Accounts     StorageAccountsClient
	BlobServices StorageBlobServicesClient
	Containers   StorageContainersClient
}

type StorageAccountsClient interface {
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
	GetBlobServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result accounts.GetServicePropertiesResult, err error)
	GetQueueServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result queues.StorageServicePropertiesResponse, err error)
}

type StorageBlobServicesClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceItems, err error)
}

type StorageContainersClient interface {
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, include storage.ListContainersInclude) (result storage.ListContainerItemsPage, err error)
}

type StorageAccountsClientImpl struct {
	*storage.AccountsClient
}

func (c StorageAccountsClientImpl) GetKeyAuthorizer(ctx context.Context, resourceGroupName string, accountName string) (*autorest.SharedKeyAuthorizer, error) {
	// use account key to create a new authorizer and then fetch service properties
	keysResult, err := c.ListKeys(ctx, resourceGroupName, accountName, "")
	if err != nil {
		return nil, err
	}
	if keysResult.Keys == nil || len(*keysResult.Keys) == 0 {
		return nil, nil
	}

	// use account key to create a new authorizer and then fetch service properties
	auth, err := autorest.NewSharedKeyAuthorizer(accountName, *(*keysResult.Keys)[0].Value, autorest.SharedKeyLite)
	if err != nil {
		return nil, nil
	}

	return auth, nil
}

func (c StorageAccountsClientImpl) GetBlobServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result accounts.GetServicePropertiesResult, err error) {
	auth, err := c.GetKeyAuthorizer(ctx, resourceGroupName, accountName)
	if err != nil {
		return accounts.GetServicePropertiesResult{}, err
	}

	blobServices := accounts.New()
	blobServices.Authorizer = auth
	return blobServices.GetServiceProperties(ctx, accountName)
}

func (c StorageAccountsClientImpl) GetQueueServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result queues.StorageServicePropertiesResponse, err error) {
	auth, err := c.GetKeyAuthorizer(ctx, resourceGroupName, accountName)
	if err != nil {
		return queues.StorageServicePropertiesResponse{}, err
	}

	queueServices := queues.New()
	queueServices.Authorizer = auth
	return queueServices.GetServiceProperties(ctx, accountName)
}

func NewStorageClient(subscriptionId string, auth autorest.Authorizer) StorageClient {
	accs := storage.NewAccountsClient(subscriptionId)
	accs.Authorizer = auth
	blobServices := storage.NewBlobServicesClient(subscriptionId)
	blobServices.Authorizer = auth
	containers := storage.NewBlobContainersClient(subscriptionId)
	containers.Authorizer = auth
	return StorageClient{
		Accounts:     StorageAccountsClientImpl{&accs},
		BlobServices: blobServices,
		Containers:   containers,
	}
}
