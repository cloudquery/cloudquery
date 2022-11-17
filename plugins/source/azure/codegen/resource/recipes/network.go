package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Network() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armnetwork.ExpressRouteCircuit),
			Resolver: &resource.FuncParams{
				Func: network.ExpressRouteCircuitsClient.NewListAllPager,
			},
		},
		{
			Struct: new(armnetwork.ExpressRouteGateway),
			Resolver: &resource.FuncParams{
				Func:   network.ExpressRouteGatewaysClient.ListBySubscription,
				Params: []string{"ctx"},
			},
		},
		{
			Struct: new(armnetwork.ExpressRoutePort),
			Resolver: &resource.FuncParams{
				Func: network.ExpressRoutePortsClient.NewListPager,
			},
		},
		{
			Struct: new(armnetwork.Interface),
			Resolver: &resource.FuncParams{
				Func: network.InterfacesClient.NewListAllPager,
			},
		},
		{
			Struct: new(armnetwork.PublicIPAddress),
			Resolver: &resource.FuncParams{
				Func: network.PublicIPAddressesClient.NewListAllPager,
			},
		},
		{
			Struct: new(armnetwork.RouteFilter),
			Resolver: &resource.FuncParams{
				Func: network.RouteFiltersClient.NewListPager,
			},
		},
		{
			Struct: new(armnetwork.RouteTable),
			Resolver: &resource.FuncParams{
				Func: network.RouteTablesClient.NewListAllPager,
			},
		},
		{
			Struct: new(armnetwork.SecurityGroup),
			Resolver: &resource.FuncParams{
				Func: network.SecurityGroupsClient.NewListAllPager,
			},
		},
		{
			Struct: new(armnetwork.VirtualNetwork),
			Resolver: &resource.FuncParams{
				Func: network.VirtualNetworksClient.NewListAllPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armnetwork.VirtualNetworkGateway),
					Resolver: &resource.FuncParams{
						Func:   network.VirtualNetworkGatewaysClient.NewListPager,
						Params: []string{"id.ResourceGroupName"},
					},
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
			Struct: new(armnetwork.Watcher),
			Resolver: &resource.FuncParams{
				Func: network.WatchersClient.NewListAllPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armnetwork.FlowLog),
					Resolver: &resource.FuncParams{
						Func:   network.FlowLogsClient.NewListPager,
						Params: []string{"id.ResourceGroupName", "*watcher.Name"},
					},
				},
			},
		},
	}
}
