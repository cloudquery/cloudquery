package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/eventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func EventHub() []*resource.Resource {
	return []*resource.Resource{
		{
			SubService: "namespaces",
			Struct:     new(armeventhub.EHNamespace),
			Resolver:   eventhub.NamespacesClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armeventhub.NetworkRuleSet),
					Resolver: eventhub.NamespacesClient.ListNetworkRuleSet,
				},
			},
		},
	}
}
