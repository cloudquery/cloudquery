package network

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNetworkExpressRouteCircuitsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	ercs := mocks.NewMockExpressRouteCircuitsClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			ExpressRouteCircuits: ercs,
		},
	}

	tid := "test"
	erca := network.ExpressRouteCircuitAuthorization{ID: &tid}
	require.Nil(t, faker.FakeData(&erca.Etag))
	require.Nil(t, faker.FakeData(&erca.Name))
	require.Nil(t, faker.FakeData(&erca.Type))
	require.Nil(t, faker.FakeData(&erca.AuthorizationPropertiesFormat))

	ercc := network.ExpressRouteCircuitConnection{ID: &tid}
	require.Nil(t, faker.FakeData(&ercc.Etag))
	require.Nil(t, faker.FakeData(&ercc.Name))
	require.Nil(t, faker.FakeData(&ercc.Type))
	require.Nil(t, faker.FakeData(&ercc.ExpressRouteCircuitConnectionPropertiesFormat))

	perc := network.PeerExpressRouteCircuitConnection{ID: &tid}
	require.Nil(t, faker.FakeData(&perc.Etag))
	require.Nil(t, faker.FakeData(&perc.Name))
	require.Nil(t, faker.FakeData(&perc.Type))
	require.Nil(t, faker.FakeData(&perc.PeerExpressRouteCircuitConnectionPropertiesFormat))

	ercp := network.ExpressRouteCircuitPeering{ID: &tid}
	require.Nil(t, faker.FakeData(&ercp.Etag))
	require.Nil(t, faker.FakeData(&ercp.Name))
	require.Nil(t, faker.FakeData(&ercp.Type))
	require.Nil(t, faker.FakeData(&ercp.ExpressRouteCircuitPeeringPropertiesFormat))
	ercp.ExpressRouteCircuitPeeringPropertiesFormat.Connections = &[]network.ExpressRouteCircuitConnection{ercc}
	ercp.ExpressRouteCircuitPeeringPropertiesFormat.PeeredConnections = &[]network.PeerExpressRouteCircuitConnection{perc}

	erc := network.ExpressRouteCircuit{ID: &tid}
	require.Nil(t, faker.FakeData(&erc.Etag))
	require.Nil(t, faker.FakeData(&erc.Location))
	require.Nil(t, faker.FakeData(&erc.Name))
	require.Nil(t, faker.FakeData(&erc.Tags))
	require.Nil(t, faker.FakeData(&erc.Type))
	require.Nil(t, faker.FakeData(&erc.Sku))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteCircuitPropertiesFormat))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteCircuitPropertiesFormat.ExpressRoutePort))
	require.Nil(t, faker.FakeData(&erc.ExpressRouteCircuitPropertiesFormat.ServiceProviderProperties))
	erc.ExpressRouteCircuitPropertiesFormat.Authorizations = &[]network.ExpressRouteCircuitAuthorization{erca}
	erc.ExpressRouteCircuitPropertiesFormat.Peerings = &[]network.ExpressRouteCircuitPeering{ercp}
	erc.ExpressRouteCircuitPropertiesFormat.ServiceProviderProvisioningState = network.ServiceProviderProvisioningState(tid)
	erc.ExpressRouteCircuitPropertiesFormat.ProvisioningState = network.ProvisioningState(tid)
	fakeId := client.FakeResourceGroup + "/" + *erc.ID
	erc.ID = &fakeId

	page := network.NewExpressRouteCircuitListResultPage(network.ExpressRouteCircuitListResult{Value: &[]network.ExpressRouteCircuit{erc}}, func(ctx context.Context, result network.ExpressRouteCircuitListResult) (network.ExpressRouteCircuitListResult, error) {
		return network.ExpressRouteCircuitListResult{}, nil
	})
	ercs.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkExpressRouteCircuits(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkExpressRouteCircuits(), buildNetworkExpressRouteCircuitsMock, client.TestOptions{})
}
