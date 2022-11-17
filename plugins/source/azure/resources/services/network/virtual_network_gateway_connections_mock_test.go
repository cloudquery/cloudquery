// Code generated by codegen; DO NOT EDIT.

package network

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/network"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildVirtualNetworkGatewayConnections(t *testing.T, ctrl *gomock.Controller, c *client.Services) {
	if c.Network == nil {
		c.Network = new(service.NetworkClient)
	}
	networkClient := c.Network
	if networkClient.VirtualNetworkGatewaysClient == nil {
		networkClient.VirtualNetworkGatewaysClient = mocks.NewMockVirtualNetworkGatewaysClient(ctrl)
	}

	mockVirtualNetworkGatewaysClient := networkClient.VirtualNetworkGatewaysClient.(*mocks.MockVirtualNetworkGatewaysClient)

	var response api.VirtualNetworkGatewaysClientListConnectionsResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockVirtualNetworkGatewaysClient.EXPECT().NewListConnectionsPager(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)
	buildVirtualNetworkGatewayConnectionsPreResolver(t, ctrl, c)
}

func buildVirtualNetworkGatewayConnectionsPreResolver(t *testing.T, ctrl *gomock.Controller, c *client.Services) {
	if c.Network == nil {
		c.Network = new(service.NetworkClient)
	}
	networkClient := c.Network
	if networkClient.VirtualNetworkGatewayConnectionsClient == nil {
		networkClient.VirtualNetworkGatewayConnectionsClient = mocks.NewMockVirtualNetworkGatewayConnectionsClient(ctrl)
	}

	mockVirtualNetworkGatewayConnectionsClient := networkClient.VirtualNetworkGatewayConnectionsClient.(*mocks.MockVirtualNetworkGatewayConnectionsClient)

	var response api.VirtualNetworkGatewayConnectionsClientGetResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.ID = to.Ptr(id)

	mockVirtualNetworkGatewayConnectionsClient.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(response, nil).MinTimes(1)
}
