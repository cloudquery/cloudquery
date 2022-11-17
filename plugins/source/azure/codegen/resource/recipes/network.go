package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Network() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armnetwork.ExpressRouteCircuit),
			Resolver: network.ExpressRouteCircuitsClient.NewListAllPager,
		},
		{
			Struct:   new(armnetwork.ExpressRouteGateway),
			Resolver: network.ExpressRouteGatewaysClient.ListBySubscription,
		},
		{
			Struct:   new(armnetwork.ExpressRoutePort),
			Resolver: network.ExpressRoutePortsClient.NewListPager,
		},
		{
			Struct:   new(armnetwork.Interface),
			Resolver: network.InterfacesClient.NewListAllPager,
		},
		{
			Struct:   new(armnetwork.PublicIPAddress),
			Resolver: network.PublicIPAddressesClient.NewListAllPager,
		},
		{
			Struct:   new(armnetwork.RouteFilter),
			Resolver: network.RouteFiltersClient.NewListPager,
		},
		{
			Struct:   new(armnetwork.RouteTable),
			Resolver: network.RouteTablesClient.NewListAllPager,
		},
		{
			Struct:   new(armnetwork.SecurityGroup),
			Resolver: network.SecurityGroupsClient.NewListAllPager,
		},
		{
			Struct:   new(armnetwork.VirtualNetwork),
			Resolver: network.VirtualNetworksClient.NewListAllPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armnetwork.VirtualNetworkGateway),
					Resolver: network.VirtualNetworkGatewaysClient.NewListPager,
					Children: []*resource.Resource{
						{
							Struct: new(armnetwork.VirtualNetworkGatewayConnection),
							Resolver: &resource.FuncParams{
								Func:   network.VirtualNetworkGatewaysClient.NewListConnectionsPager,
								Params: []string{"id.ResourceGroupName", "*virtualNetworkGateway.Name"},
							},
							PreResolver: &resource.FuncParams{
								Func:       network.VirtualNetworkGatewayConnectionsClient.Get,
								Params:     []string{"ctx", "id.ResourceGroupName", "*virtualNetworkGatewayConnectionListEntity.Name"},
								BasicValue: new(armnetwork.VirtualNetworkGatewayConnectionListEntity),
							},
						},
					},
				},
			},
		},
		{
			Struct:   new(armnetwork.Watcher),
			Resolver: network.WatchersClient.NewListAllPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armnetwork.FlowLog),
					Resolver: network.FlowLogsClient.NewListPager,
				},
			},
		},
	}
}
