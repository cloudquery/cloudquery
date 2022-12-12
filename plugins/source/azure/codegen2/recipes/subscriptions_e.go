package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"

func ArmSubscriptionsE() []Table {
	tables := []Table{
		{
			Service:        "armsubscription",
			Name:           "subscriptions",
			Struct:         &armsubscription.Subscription{},
			ResponseStruct: &armsubscription.SubscriptionsClientListResponse{},
			Client:         &armsubscription.SubscriptionsClient{},
			ListFunc:       (&armsubscription.SubscriptionsClient{}).NewListPager,
			NewFunc:        armsubscription.NewSubscriptionsClient,
			URL:            "/subscriptions",
			Multiplex:      `client.SingleSubscriptionMultiplex`,
			SkipFetch:      true,
		},
		{
			Service:        "armsubscription",
			Name:           "tenants",
			Struct:         &armsubscription.TenantIDDescription{},
			ResponseStruct: &armsubscription.TenantsClientListResponse{},
			Client:         &armsubscription.TenantsClient{},
			ListFunc:       (&armsubscription.TenantsClient{}).NewListPager,
			NewFunc:        armsubscription.NewTenantsClient,
			URL:            "/tenants",
			Multiplex:      `client.SingleSubscriptionMultiplex`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, ArmSubscriptionsE()...)
}
