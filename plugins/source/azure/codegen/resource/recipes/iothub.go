package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/iothub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func IoTHub() []*resource.Resource {
	return []*resource.Resource{
		{
			Name:       "azure_iot_hubs",
			SubService: "hubs",
			Struct:     new(armiothub.Description),
			Resolver: &resource.FuncParams{
				Func: iothub.ResourceClient.NewListBySubscriptionPager,
			},
		},
	}
}
