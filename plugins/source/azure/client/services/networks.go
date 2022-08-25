//go:generate mockgen -destination=./mocks/network.go -package=mocks . NetworkExpressRouteCircuitsClient,NetworkExpressRouteGatewaysClient,NetworkExpressRoutePortsClient,NetworkInterfacesClient,NetworkPublicIPAddressesClient,NetworkRouteFiltersClient,NetworkRouteTablesClient,NetworkSecurityGroupsClient,NetworkVirtualNetworkGatewaysClient,NetworkVirtualNetworksClient,NetworkWatchersClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/Azure/go-autorest/autorest"
)

type NetworkClient struct {
	ExpressRouteCircuits   NetworkExpressRouteCircuitsClient
	ExpressRouteGateways   NetworkExpressRouteGatewaysClient
	ExpressRoutePorts      NetworkExpressRoutePortsClient
	Interfaces             NetworkInterfacesClient
	PublicIPAddresses      NetworkPublicIPAddressesClient
	RouteFilters           NetworkRouteFiltersClient
	RouteTables            NetworkRouteTablesClient
	SecurityGroups         NetworkSecurityGroupsClient
	VirtualNetworkGateways NetworkVirtualNetworkGatewaysClient
	VirtualNetworks        NetworkVirtualNetworksClient
	Watchers               NetworkWatchersClient
}
type NetworkExpressRouteCircuitsClient interface {
	ListAll(ctx context.Context) (result network.ExpressRouteCircuitListResultPage, err error)
}

type NetworkExpressRouteGatewaysClient interface {
	ListBySubscription(ctx context.Context) (result network.ExpressRouteGatewayList, err error)
}

type NetworkExpressRoutePortsClient interface {
	List(ctx context.Context) (result network.ExpressRoutePortListResultPage, err error)
}

type NetworkInterfacesClient interface {
	ListAll(ctx context.Context) (result network.InterfaceListResultPage, err error)
}

type NetworkPublicIPAddressesClient interface {
	ListAll(ctx context.Context) (result network.PublicIPAddressListResultPage, err error)
}

type NetworkRouteFiltersClient interface {
	List(ctx context.Context) (result network.RouteFilterListResultPage, err error)
}

type NetworkRouteTablesClient interface {
	ListAll(ctx context.Context) (result network.RouteTableListResultPage, err error)
}

type NetworkSecurityGroupsClient interface {
	ListAll(ctx context.Context) (result network.SecurityGroupListResultPage, err error)
}

type NetworkVirtualNetworkGatewaysClient interface {
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayListResultPage, err error)
	ListConnections(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewayListConnectionsResultPage, err error)
}

type NetworkVirtualNetworksClient interface {
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
}

type NetworkWatchersClient interface {
	ListAll(ctx context.Context) (result network.WatcherListResult, err error)
	GetFlowLogStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogStatusParameters) (result network.WatchersGetFlowLogStatusFuture, err error)
}

func NewNetworksClient(subscriptionId string, auth autorest.Authorizer) NetworkClient {
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
	return NetworkClient{
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
