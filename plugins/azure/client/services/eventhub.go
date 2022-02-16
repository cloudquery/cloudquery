package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	"github.com/Azure/go-autorest/autorest"
)

type EventHubClient interface {
	List(ctx context.Context) (result eventhub.EHNamespaceListResultPage, err error)
}

func NewEventHubClient(subscriptionId string, auth autorest.Authorizer) EventHubClient {
	cl := eventhub.NewNamespacesClient(subscriptionId)
	cl.Authorizer = auth
	return cl
}
