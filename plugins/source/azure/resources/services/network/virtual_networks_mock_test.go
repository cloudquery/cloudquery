// Auto generated code - DO NOT EDIT.

package network

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func TestNetworkVirtualNetworks(t *testing.T) {
	client.MockTestHelper(t, VirtualNetworks(), createVirtualNetworksMock)
}

func createVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworksClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworks:        mockClient,
			VirtualNetworkGateways: createVirtualNetworkGatewaysMock(t, ctrl).Network.VirtualNetworkGateways,
		},
	}

	data := network.VirtualNetwork{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{data}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
