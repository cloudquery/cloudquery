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

func TestNetworkRouteTables(t *testing.T) {
	client.MockTestHelper(t, RouteTables(), createRouteTablesMock)
}

func createRouteTablesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkRouteTablesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			RouteTables: mockClient,
		},
	}

	data := network.RouteTable{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := network.NewRouteTableListResultPage(network.RouteTableListResult{Value: &[]network.RouteTable{data}}, func(ctx context.Context, result network.RouteTableListResult) (network.RouteTableListResult, error) {
		return network.RouteTableListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
