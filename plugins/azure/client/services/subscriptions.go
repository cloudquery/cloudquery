package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest"
)

type SubscriptionsClient struct {
	SubscriptionID string
	Subscriptions  SubscriptionGetter
}

func NewSubscriptionsClient(subscriptionId string, auth autorest.Authorizer) SubscriptionsClient {
	s := subscription.NewSubscriptionsClient()
	s.Authorizer = auth
	return SubscriptionsClient{
		SubscriptionID: subscriptionId,
		Subscriptions:  s,
	}
}

type SubscriptionGetter interface {
	Get(ctx context.Context, subscriptionID string) (result subscription.Model, err error)
}
