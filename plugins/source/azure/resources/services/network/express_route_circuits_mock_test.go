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

func buildExpressRouteCircuits(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockExpressRouteCircuitsClient := mocks.NewMockExpressRouteCircuitsClient(ctrl)

	var response api.ExpressRouteCircuitsClientListAllResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockExpressRouteCircuitsClient.EXPECT().NewListAllPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	networkClient := &service.NetworkClient{
		ExpressRouteCircuitsClient: mockExpressRouteCircuitsClient,
	}

	c := &client.Services{Network: networkClient}

	return c
}

func TestExpressRouteCircuits(t *testing.T) {
	client.MockTestHelper(t, ExpressRouteCircuits(), buildExpressRouteCircuits)
}
