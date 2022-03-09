package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	"github.com/Azure/go-autorest/autorest"
)

type BatchAccountClient interface {
	List(ctx context.Context) (result batch.AccountListResultPage, err error)
}

type BatchClient struct {
	Account BatchAccountClient
}

func NewBatchClient(subscriptionId string, auth autorest.Authorizer) BatchClient {
	acCl := batch.NewAccountClient(subscriptionId)
	acCl.Authorizer = auth
	return BatchClient{
		Account: acCl,
	}
}
