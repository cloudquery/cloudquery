//go:generate mockgen -destination=./mocks/eventhub.go -package=mocks . EventHubNamespacesClient,EventHubNetworkRuleSetsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	"github.com/Azure/go-autorest/autorest"
)

type EventHubClient struct {
	Namespaces      EventHubNamespacesClient
	NetworkRuleSets EventHubNetworkRuleSetsClient
}

type EventHubNamespacesClient interface {
	List(ctx context.Context) (result eventhub.EHNamespaceListResultPage, err error)
}

type EventHubNetworkRuleSetsClient interface {
	GetNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NetworkRuleSet, err error)
}

func NewEventHubClient(subscriptionId string, auth autorest.Authorizer) EventHubClient {
	cl := eventhub.NewNamespacesClient(subscriptionId)
	cl.Authorizer = auth
	return EventHubClient{Namespaces: cl, NetworkRuleSets: cl}
}
