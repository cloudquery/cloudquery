package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

func ArmResourcesE() []Table {
	tables := []Table{
		{
			Service:        "armresources",
			Name:           "resources",
			Struct:         &armresources.GenericResourceExpanded{},
			ResponseStruct: &armresources.ClientListResponse{},
			Client:         &armresources.Client{},
			ListFunc:       (&armresources.Client{}).NewListPager,
			NewFunc:        armresources.NewClient,
			URL:            "/subscriptions/{subscriptionId}/resources",
			Multiplex:      `client.SubscriptionMultiplex`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, ArmResourcesE()...)
}
