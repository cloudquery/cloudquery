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

func buildNetworkExpressRoutePortsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	rpc := mocks.NewMockExpressRoutePortsClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			ExpressRoutePorts: rpc,
		},
	}

	tid := "test"
	erl := network.ExpressRouteLink{ID: &tid}
	require.Nil(t, faker.FakeData(&erl.Etag))
	require.Nil(t, faker.FakeData(&erl.Name))
	require.Nil(t, faker.FakeData(&erl.ExpressRouteLinkPropertiesFormat))
	require.Nil(t, faker.FakeData(&erl.ExpressRouteLinkPropertiesFormat.MacSecConfig))

	erc := network.ExpressRoutePort{ID: &tid}
	require.Nil(t, faker.FakeData(&erc.Etag))
	require.Nil(t, faker.FakeData(&erc.Location))
	require.Nil(t, faker.FakeData(&erc.Name))
	require.Nil(t, faker.FakeData(&erc.Tags))
	require.Nil(t, faker.FakeData(&erc.Type))
	require.Nil(t, faker.FakeData(&erc.Identity))
	require.Nil(t, faker.FakeData(&erc.ExpressRoutePortPropertiesFormat))
	erc.ExpressRoutePortPropertiesFormat.Links = &[]network.ExpressRouteLink{erl}

	page := network.NewExpressRoutePortListResultPage(network.ExpressRoutePortListResult{Value: &[]network.ExpressRoutePort{erc}}, func(ctx context.Context, result network.ExpressRoutePortListResult) (network.ExpressRoutePortListResult, error) {
		return network.ExpressRoutePortListResult{}, nil
	})
	rpc.EXPECT().List(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkExpressRoutePorts(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkExpressRoutePorts(), buildNetworkExpressRoutePortsMock, client.TestOptions{})
}
