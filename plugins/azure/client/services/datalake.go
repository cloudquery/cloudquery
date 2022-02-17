//go:generate mockgen -destination=./mocks/datalake.go -package=mocks . DataLakeStorageAccountsClient,DataLakeAnalyticsAccountsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account"
	storeAccount "github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
	"github.com/Azure/go-autorest/autorest"
)

type DataLakeClient struct {
	DataLakeStorageAccounts   DataLakeStorageAccountsClient
	DataLakeAnalyticsAccounts DataLakeAnalyticsAccountsClient
}

func NewDataLakeClient(subscriptionId string, auth autorest.Authorizer) DataLakeClient {
	storeAccounts := storeAccount.NewAccountsClient(subscriptionId)
	storeAccounts.Authorizer = auth
	analyticsAccounts := account.NewAccountsClient(subscriptionId)
	analyticsAccounts.Authorizer = auth
	return DataLakeClient{
		DataLakeStorageAccounts:   storeAccounts,
		DataLakeAnalyticsAccounts: analyticsAccounts,
	}
}

type DataLakeStorageAccountsClient interface {
	List(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result storeAccount.DataLakeStoreAccountListResultPage, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result storeAccount.DataLakeStoreAccount, err error)
}

type DataLakeAnalyticsAccountsClient interface {
	List(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeAnalyticsAccount, err error)
}
