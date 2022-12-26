package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"

func init() {
	tables := []Table{
		{
			Service:        "armsearch",
			Name:           "services",
			Struct:         &armsearch.Service{},
			ResponseStruct: &armsearch.ServicesClientListBySubscriptionResponse{},
			Client:         &armsearch.ServicesClient{},
			ListFunc:       (&armsearch.ServicesClient{}).NewListBySubscriptionPager,
			NewFunc:        armsearch.NewServicesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Search/searchServices",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_search)`,
			SkipFetch:      true,
		},
	}
	Tables = append(Tables, tables...)
}
