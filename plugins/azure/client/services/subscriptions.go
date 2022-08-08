//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionsClient,TenantsClient
package services

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Azure/go-autorest/autorest"
)

type Subscriptions struct {
	SubscriptionID string
	Subscriptions  SubscriptionsClient
	Tenants        TenantsClient
}

type SubscriptionsClient interface {
	NewListPager(options *armsubscriptions.ClientListOptions) *runtime.Pager[armsubscriptions.ClientListResponse]
	NewListLocationsPager(subscriptionId string, options *armsubscriptions.ClientListLocationsOptions) *runtime.Pager[armsubscriptions.ClientListLocationsResponse]
}

type TenantsClient interface {
	NewListPager(options *armsubscriptions.TenantsClientListOptions) *runtime.Pager[armsubscriptions.TenantsClientListResponse]
}

func NewSubscriptionsClient(subscriptionId string, auth autorest.Authorizer, azCred azcore.TokenCredential) (Subscriptions, error) {
	s, err := armsubscriptions.NewClient(azCred, nil)
	if err != nil {
		return Subscriptions{}, fmt.Errorf("failed to create subscriptions client: %w", err)
	}

	t, err := armsubscriptions.NewTenantsClient(azCred, nil)
	if err != nil {
		return Subscriptions{}, fmt.Errorf("failed to create tenants client: %w", err)
	}

	return Subscriptions{
		SubscriptionID: subscriptionId,
		Subscriptions:  s,
		Tenants:        t,
	}, nil
}
