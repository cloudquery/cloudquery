// Code generated by codegen; DO NOT EDIT.
package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

type (
	RuntimePagerArmnetworkVirtualNetworkGatewayConnectionsClientListResponse = runtime.Pager[armnetwork.VirtualNetworkGatewayConnectionsClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/network/virtual_network_gateway_connections.go -source=virtual_network_gateway_connections.go VirtualNetworkGatewayConnectionsClient
type VirtualNetworkGatewayConnectionsClient interface {
	Get(context.Context, string, string, *armnetwork.VirtualNetworkGatewayConnectionsClientGetOptions) (armnetwork.VirtualNetworkGatewayConnectionsClientGetResponse, error)
	GetSharedKey(context.Context, string, string, *armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyOptions) (armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse, error)
	NewListPager(string, *armnetwork.VirtualNetworkGatewayConnectionsClientListOptions) *RuntimePagerArmnetworkVirtualNetworkGatewayConnectionsClientListResponse
}
