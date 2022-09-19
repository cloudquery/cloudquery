//go:generate mockgen -destination=./mocks/batch.go -package=mocks . BatchAccountsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	"github.com/Azure/go-autorest/autorest"
)

type BatchClient struct {
	Accounts BatchAccountsClient
}

type BatchAccountsClient interface {
	List(ctx context.Context) (result batch.AccountListResultPage, err error)
}

func NewBatchClient(subscriptionId string, auth autorest.Authorizer) BatchClient {
	acCl := batch.NewAccountClient(subscriptionId)
	acCl.Authorizer = auth
	return BatchClient{
		Accounts: acCl,
	}
}
