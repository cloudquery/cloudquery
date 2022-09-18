// Auto generated code - DO NOT EDIT.

package network

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func TestNetworkExpressRouteGateways(t *testing.T) {
	client.MockTestHelper(t, ExpressRouteGateways(), createExpressRouteGatewaysMock)
}

func createExpressRouteGatewaysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRouteGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRouteGateways: mockClient,
		},
	}

	data := network.ExpressRouteGateway{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := network.ExpressRouteGatewayList{Value: &[]network.ExpressRouteGateway{data}}

	mockClient.EXPECT().ListBySubscription(gomock.Any()).Return(result, nil)
	return s
}
