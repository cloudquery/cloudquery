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

func createVirtualNetworkGatewaysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworkGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworkGateways:           mockClient,
			VirtualNetworkGatewayConnections: createVirtualNetworkGatewayConnectionsMock(t, ctrl).Network.VirtualNetworkGatewayConnections,
		},
	}

	data := network.VirtualNetworkGateway{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := network.NewVirtualNetworkGatewayListResultPage(network.VirtualNetworkGatewayListResult{Value: &[]network.VirtualNetworkGateway{data}}, func(ctx context.Context, result network.VirtualNetworkGatewayListResult) (network.VirtualNetworkGatewayListResult, error) {
		return network.VirtualNetworkGatewayListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test").Return(result, nil)
	return s
}
