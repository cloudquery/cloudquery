package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/redis"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Redis() []*resource.Resource {
	return []*resource.Resource{
		{
			SubService: "caches",
			Struct:     new(armredis.ResourceInfo),
			Resolver: &resource.FuncParams{
				Func: redis.Client.NewListBySubscriptionPager,
			},
		},
	}
}
