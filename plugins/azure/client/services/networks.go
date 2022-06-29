package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/Azure/go-autorest/autorest"
)

type NetworksClient struct {
	ExpressRouteCircuits   ExpressRouteCircuitsClient
	ExpressRouteGateways   ExpressRouteGatewaysClient
	ExpressRoutePorts      ExpressRoutePortsClient
	Interfaces             InterfacesClient
	PublicIPAddresses      PublicIPAddressesClient
	RouteFilters           RouteFiltersClient
	RouteTables            RouteTablesClient
	SecurityGroups         SecurityGroupsClient
	VirtualNetworkGateways VirtualNetworkGatewaysClient
	VirtualNetworks        VirtualNetworksClient
	Watchers               WatchersClient
}
type ExpressRouteCircuitsClient interface {
	ListAll(ctx context.Context) (result network.ExpressRouteCircuitListResultPage, err error)
}

type ExpressRouteGatewaysClient interface {
	ListBySubscription(ctx context.Context) (result network.ExpressRouteGatewayList, err error)
}

type ExpressRoutePortsClient interface {
	List(ctx context.Context) (result network.ExpressRoutePortListResultPage, err error)
}

type InterfacesClient interface {
	ListAll(ctx context.Context) (result network.InterfaceListResultPage, err error)
}

type PublicIPAddressesClient interface {
	ListAll(ctx context.Context) (result network.PublicIPAddressListResultPage, err error)
}

type RouteFiltersClient interface {
	List(ctx context.Context) (result network.RouteFilterListResultPage, err error)
}

type RouteTablesClient interface {
	ListAll(ctx context.Context) (result network.RouteTableListResultPage, err error)
}

type SecurityGroupsClient interface {
	ListAll(ctx context.Context) (result network.SecurityGroupListResultPage, err error)
}

type VirtualNetworkGatewaysClient interface {
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayListResultPage, err error)
	ListConnections(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewayListConnectionsResultPage, err error)
}

type VirtualNetworksClient interface {
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
}

type WatchersClient interface {
	ListAll(ctx context.Context) (result network.WatcherListResult, err error)
	GetFlowLogStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogStatusParameters) (result network.WatchersGetFlowLogStatusFuture, err error)
}

func NewNetworksClient(subscriptionId string, auth autorest.Authorizer) NetworksClient {
	erc := network.NewExpressRouteCircuitsClient(subscriptionId)
	erc.Authorizer = auth
	erg := network.NewExpressRouteGatewaysClient(subscriptionId)
	erg.Authorizer = auth
	erp := network.NewExpressRoutePortsClient(subscriptionId)
	erp.Authorizer = auth
	ifs := network.NewInterfacesClient(subscriptionId)
	ifs.Authorizer = auth
	pips := network.NewPublicIPAddressesClient(subscriptionId)
	pips.Authorizer = auth
	rf := network.NewRouteFiltersClient(subscriptionId)
	rf.Authorizer = auth
	rt := network.NewRouteTablesClient(subscriptionId)
	rt.Authorizer = auth
	sg := network.NewSecurityGroupsClient(subscriptionId)
	sg.Authorizer = auth
	vng := network.NewVirtualNetworkGatewaysClient(subscriptionId)
	vng.Authorizer = auth
	vn := network.NewVirtualNetworksClient(subscriptionId)
	vn.Authorizer = auth
	wch := network.NewWatchersClient(subscriptionId)
	wch.Authorizer = auth
	return NetworksClient{
		ExpressRouteCircuits:   erc,
		ExpressRouteGateways:   erg,
		ExpressRoutePorts:      erp,
		Interfaces:             ifs,
		PublicIPAddresses:      pips,
		RouteFilters:           rf,
		RouteTables:            rt,
		SecurityGroups:         sg,
		VirtualNetworkGateways: vng,
		VirtualNetworks:        vn,
		Watchers:               wch,
	}
}
