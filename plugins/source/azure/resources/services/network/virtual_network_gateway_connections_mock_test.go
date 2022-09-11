// Auto generated code - DO NOT EDIT.

package network

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func createVirtualNetworkGatewayConnectionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworkGatewayConnectionsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworkGatewayConnections: mockClient,
		},
	}

	data := network.VirtualNetworkGatewayConnection{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewVirtualNetworkGatewayConnectionListResultPage(network.VirtualNetworkGatewayConnectionListResult{Value: &[]network.VirtualNetworkGatewayConnection{data}}, func(ctx context.Context, result network.VirtualNetworkGatewayConnectionListResult) (network.VirtualNetworkGatewayConnectionListResult, error) {
		return network.VirtualNetworkGatewayConnectionListResult{}, nil
	})

	mockClient.EXPECT().ListConnections(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
