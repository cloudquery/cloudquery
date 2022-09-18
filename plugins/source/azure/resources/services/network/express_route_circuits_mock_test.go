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

func TestNetworkExpressRouteCircuits(t *testing.T) {
	client.MockTestHelper(t, ExpressRouteCircuits(), createExpressRouteCircuitsMock)
}

func createExpressRouteCircuitsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRouteCircuitsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRouteCircuits: mockClient,
		},
	}

	data := network.ExpressRouteCircuit{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewExpressRouteCircuitListResultPage(network.ExpressRouteCircuitListResult{Value: &[]network.ExpressRouteCircuit{data}}, func(ctx context.Context, result network.ExpressRouteCircuitListResult) (network.ExpressRouteCircuitListResult, error) {
		return network.ExpressRouteCircuitListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
