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

func TestNetworkExpressRoutePorts(t *testing.T) {
	client.MockTestHelper(t, ExpressRoutePorts(), createExpressRoutePortsMock)
}

func createExpressRoutePortsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRoutePortsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRoutePorts: mockClient,
		},
	}

	data := network.ExpressRoutePort{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := network.NewExpressRoutePortListResultPage(network.ExpressRoutePortListResult{Value: &[]network.ExpressRoutePort{data}}, func(ctx context.Context, result network.ExpressRoutePortListResult) (network.ExpressRoutePortListResult, error) {
		return network.ExpressRoutePortListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
