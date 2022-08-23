//go:generate mockgen -destination=./mocks/batch.go -package=mocks . AccountsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	"github.com/Azure/go-autorest/autorest"
)

type AccountsClient interface {
	List(ctx context.Context) (result batch.AccountListResultPage, err error)
}

type BatchClient struct {
	Accounts AccountsClient
}

func NewBatchClient(subscriptionId string, auth autorest.Authorizer) BatchClient {
	acCl := batch.NewAccountClient(subscriptionId)
	acCl.Authorizer = auth
	return BatchClient{
		Accounts: acCl,
	}
}
