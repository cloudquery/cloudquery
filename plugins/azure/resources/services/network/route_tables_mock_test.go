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

func fakeRouteTableRoute(t *testing.T) network.Route {
	var r network.Route
	require.NoError(t, faker.FakeData(&r))
	fakeId := client.FakeResourceGroup + "/" + *r.ID
	r.ID = &fakeId
	return r
}

func fakeRouteTableSubnet(t *testing.T) network.Subnet {
	var s network.Subnet
	require.NoError(t, faker.FakeDataSkipFields(&s, []string{"SubnetPropertiesFormat"}))
	fakeId := client.FakeResourceGroup + "/" + *s.ID
	s.ID = &fakeId
	return s
}

func fakeRouteTable(t *testing.T) network.RouteTable {
	rt := network.RouteTable{
		RouteTablePropertiesFormat: &network.RouteTablePropertiesFormat{
			Routes: &[]network.Route{
				fakeRouteTableRoute(t),
			},
			Subnets: &[]network.Subnet{
				fakeRouteTableSubnet(t),
			},
		},
	}
	require.NoError(t, faker.FakeDataSkipFields(&rt, []string{"RouteTablePropertiesFormat"}))
	require.NoError(t, faker.FakeDataSkipFields(rt.RouteTablePropertiesFormat, []string{"Routes", "Subnets", "ProvisioningState"}))

	fakeId := client.FakeResourceGroup + "/" + *rt.ID
	rt.ID = &fakeId
	return rt
}

func buildNetworkRouteTablesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	rtc := mocks.NewMockRouteTablesClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			RouteTables: rtc,
		},
	}
	rt := fakeRouteTable(t)
	rtp := network.NewRouteTableListResultPage(network.RouteTableListResult{Value: &[]network.RouteTable{rt}}, func(ctx context.Context, result network.RouteTableListResult) (network.RouteTableListResult, error) {
		return network.RouteTableListResult{}, nil
	})
	rtc.EXPECT().ListAll(gomock.Any()).Return(rtp, nil)
	return s
}

func TestNetworkRouteTables(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkRouteTables(), buildNetworkRouteTablesMock, client.TestOptions{})
}
