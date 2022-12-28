package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"

func init() {
	tables := []Table{
		{
			Service:        "armredis",
			Name:           "caches",
			Struct:         &armredis.ResourceInfo{},
			ResponseStruct: &armredis.ClientListBySubscriptionResponse{},
			Client:         &armredis.Client{},
			ListFunc:       (&armredis.Client{}).NewListBySubscriptionPager,
			NewFunc:        armredis.NewClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cache/redis",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_cache)`,
		},
	}
	Tables = append(Tables, tables...)
}
