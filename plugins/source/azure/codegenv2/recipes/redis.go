package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

func RedisResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "caches",
			Struct: &armredis.ResourceInfo{},
			ResponseStruct: &armredis.ClientListBySubscriptionResponse{},
			Client: &armredis.Client{},
			ListFunc: (&armredis.Client{}).NewListBySubscriptionPager,
			NewFunc: armredis.NewClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "redis/armredis"
		r.Service = "armredis"
		r.Template = "list"
	}

	return resources
}