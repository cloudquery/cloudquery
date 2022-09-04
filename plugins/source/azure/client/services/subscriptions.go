//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionsSubscriptionsClient,SubscriptionsTenantsClient,SubscriptionsLocationsClient
package services

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Azure/go-autorest/autorest"
)

type SubscriptionsClient struct {
	SubscriptionID string
	Subscriptions  SubscriptionsSubscriptionsClient
	Locations      SubscriptionsLocationsClient
	Tenants        SubscriptionsTenantsClient
}

type SubscriptionsSubscriptionsClient interface {
	NewListPager(options *armsubscriptions.ClientListOptions) *runtime.Pager[armsubscriptions.ClientListResponse]
}

type SubscriptionsLocationsClient interface {
	NewListLocationsPager(subscriptionId string, options *armsubscriptions.ClientListLocationsOptions) *runtime.Pager[armsubscriptions.ClientListLocationsResponse]
}

type SubscriptionsTenantsClient interface {
	NewListPager(options *armsubscriptions.TenantsClientListOptions) *runtime.Pager[armsubscriptions.TenantsClientListResponse]
}

func NewSubscriptionsClient(subscriptionId string, auth autorest.Authorizer, azCred azcore.TokenCredential) (SubscriptionsClient, error) {
	s, err := armsubscriptions.NewClient(azCred, nil)
	if err != nil {
		return SubscriptionsClient{}, fmt.Errorf("failed to create subscriptions client: %w", err)
	}

	t, err := armsubscriptions.NewTenantsClient(azCred, nil)
	if err != nil {
		return SubscriptionsClient{}, fmt.Errorf("failed to create tenants client: %w", err)
	}

	return SubscriptionsClient{
		SubscriptionID: subscriptionId,
		Subscriptions:  s,
		Tenants:        t,
		Locations:      s,
	}, nil
}
