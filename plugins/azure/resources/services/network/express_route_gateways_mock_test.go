package network

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNetworkExpressRouteGatewaysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	rgc := mocks.NewMockExpressRouteGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			ExpressRouteGateways: rgc,
		},
	}

	tid := "test"
	erc := network.ExpressRouteConnection{ID: &tid}
	require.Nil(t, faker.FakeData(&erc.Name))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties.ExpressRouteCircuitPeering))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties.ProvisioningState))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties.RoutingConfiguration))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties.RoutingConfiguration.AssociatedRouteTable.ID))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties.RoutingConfiguration.PropagatedRouteTables))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteConnectionProperties.RoutingConfiguration.VnetRoutes))

	erg := network.ExpressRouteGateway{ID: &tid}
	require.Nil(t, faker.FakeData(&erg.Etag))
	require.Nil(t, faker.FakeData(&erg.Location))
	require.Nil(t, faker.FakeData(&erg.Name))
	require.Nil(t, faker.FakeData(&erg.Tags))
	require.Nil(t, faker.FakeData(&erg.Type))
	require.Nil(t, faker.FakeData(&erg.ExpressRouteGatewayProperties))
	require.Nil(t, faker.FakeData(&erg.ExpressRouteGatewayProperties.AutoScaleConfiguration))
	require.Nil(t, faker.FakeData(&erg.ExpressRouteGatewayProperties.AutoScaleConfiguration.Bounds))
	require.Nil(t, faker.FakeData(&erg.ExpressRouteGatewayProperties.ProvisioningState))
	require.Nil(t, faker.FakeData(&erg.ExpressRouteGatewayProperties.VirtualHub))
	erg.ExpressRouteGatewayProperties.ExpressRouteConnections = &[]network.ExpressRouteConnection{erc}
	fakeId := client.FakeResourceGroup + "/" + *erg.ID
	erg.ID = &fakeId

	list := network.ExpressRouteGatewayList{Value: &[]network.ExpressRouteGateway{erg}}
	rgc.EXPECT().ListBySubscription(gomock.Any()).Return(list, nil)
	return s
}

func TestNetworkExpressRouteGateways(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkExpressRouteGateways(), buildNetworkExpressRouteGatewaysMock, client.TestOptions{})
}
