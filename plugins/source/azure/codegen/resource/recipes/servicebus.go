package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func ServiceBus() []*resource.Resource {
	return []*resource.Resource{
		{
			SubService: "namespaces",
			Struct:     new(armservicebus.SBNamespace),
			Resolver:   servicebus.NamespacesClient.NewListPager,
			Children: []*resource.Resource{
				{
					SubService: "topics",
					Struct:     new(armservicebus.SBTopic),
					Resolver:   servicebus.TopicsClient.NewListByNamespacePager,
					Children: []*resource.Resource{
						{
							SubService: "authorization_rules",
							Struct:     new(armservicebus.SBAuthorizationRule),
							Resolver:   servicebus.TopicsClient.NewListAuthorizationRulesPager,
							Children: []*resource.Resource{
								{
									Struct:   new(armservicebus.AccessKeys),
									Resolver: servicebus.TopicsClient.ListKeys,
								},
							},
						},
					},
				},
			},
		},
	}
}
