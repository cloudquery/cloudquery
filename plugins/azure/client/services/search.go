package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
)

type SearchClient struct {
	Service SearchServiceClient
}

type SearchServiceClient interface {
	ListBySubscription(ctx context.Context, clientRequestID *uuid.UUID) (result search.ServiceListResultPage, err error)
}

func NewSearchClient(subscriptionId string, auth autorest.Authorizer) SearchClient {
	cl := search.NewServicesClient(subscriptionId)
	cl.Authorizer = auth

	return SearchClient{
		Service: cl,
	}
}
