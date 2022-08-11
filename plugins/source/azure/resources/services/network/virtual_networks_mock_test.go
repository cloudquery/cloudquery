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

func fakeSubnet(t *testing.T) network.Subnet {
	var subnet network.Subnet
	require.NoError(t, faker.FakeDataSkipFields(&subnet, []string{"SubnetPropertiesFormat"}))
	var subnetPropertiesFormat network.SubnetPropertiesFormat
	require.NoError(t, faker.FakeDataSkipFields(&subnetPropertiesFormat, []string{
		"NetworkSecurityGroup",
		"RouteTable",
		"ServiceEndpointPolicies",
		"PrivateEndpoints",
		"IPConfigurations",
		"IPConfigurationProfiles",

		"ProvisioningState",
		"PrivateEndpointNetworkPolicies",
		"PrivateLinkServiceNetworkPolicies",
	}))
	guid := "guid"
	subnetPropertiesFormat.ProvisioningState = network.ProvisioningStateSucceeded
	subnetPropertiesFormat.PrivateEndpointNetworkPolicies = network.VirtualNetworkPrivateEndpointNetworkPoliciesDisabled
	subnetPropertiesFormat.PrivateLinkServiceNetworkPolicies = network.VirtualNetworkPrivateLinkServiceNetworkPoliciesDisabled
	nsg := network.SecurityGroup{SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
		ResourceGUID:      &guid,
		ProvisioningState: network.ProvisioningStateDeleting,
	}}
	require.NoError(t, faker.FakeDataSkipFields(&nsg, []string{"SecurityGroupPropertiesFormat"}))
	subnetPropertiesFormat.NetworkSecurityGroup = &nsg
	var rt network.RouteTable
	require.NoError(t, faker.FakeDataSkipFields(&rt, []string{"RouteTablePropertiesFormat"}))
	var b bool
	rt.RouteTablePropertiesFormat = &network.RouteTablePropertiesFormat{
		DisableBgpRoutePropagation: &b,
		ResourceGUID:               &guid,
		ProvisioningState:          network.ProvisioningStateDeleting,
	}
	subnetPropertiesFormat.RouteTable = &rt
	var ipconfig network.IPConfiguration
	require.NoError(t, faker.FakeDataSkipFields(&ipconfig, []string{"IPConfigurationPropertiesFormat"}))
	subnetPropertiesFormat.IPConfigurations = &[]network.IPConfiguration{ipconfig}

	var pe network.PrivateEndpoint
	require.NoError(t, faker.FakeDataSkipFields(&pe, []string{"PrivateEndpointProperties"}))
	subnetPropertiesFormat.PrivateEndpoints = &[]network.PrivateEndpoint{pe}

	subnet.SubnetPropertiesFormat = &subnetPropertiesFormat
	return subnet
}

func fakeVirtualNetwork(t *testing.T) network.VirtualNetwork {
	vn := network.VirtualNetwork{
		VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
			Subnets: &[]network.Subnet{
				fakeSubnet(t),
			},
		},
	}
	require.NoError(t, faker.FakeDataSkipFields(&vn, []string{"VirtualNetworkPropertiesFormat"}))
	require.NoError(t, faker.FakeDataSkipFields(vn.VirtualNetworkPropertiesFormat, []string{"Subnets", "ProvisioningState"}))

	fakeId := client.FakeResourceGroup + "/" + *vn.ID
	vn.ID = &fakeId
	vn.DhcpOptions.DNSServers = &[]string{faker.IPv4()}
	return vn
}

func fakeVirtualNetworkGateway(t *testing.T) network.VirtualNetworkGateway {
	vng := network.VirtualNetworkGateway{}
	require.NoError(t, faker.FakeData(&vng))
	fakeVngID := client.FakeResourceGroup + "/" + *vng.ID
	vng.ID = &fakeVngID
	return vng
}

func fakeVirtualNetworkGatewayConnection(t *testing.T) network.VirtualNetworkGatewayConnectionListEntity {
	vngc := network.VirtualNetworkGatewayConnectionListEntity{}
	require.NoError(t, faker.FakeDataSkipFields(&vngc, []string{"VirtualNetworkGatewayConnectionListEntityPropertiesFormat"}))
	require.NoError(t, faker.FakeData(&vngc.VirtualNetworkGatewayConnectionListEntityPropertiesFormat))
	fakeVngcID := client.FakeResourceGroup + "/" + *vngc.ID
	vngc.ID = &fakeVngcID
	return vngc
}

func buildNetworkVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	n := mocks.NewMockVirtualNetworksClient(ctrl)
	ngc := mocks.NewMockVirtualNetworkGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			VirtualNetworks:        n,
			VirtualNetworkGateways: ngc,
		},
	}

	vn := fakeVirtualNetwork(t)
	vng := fakeVirtualNetworkGateway(t)
	vngc := fakeVirtualNetworkGatewayConnection(t)

	vnp := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{vn}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})
	vngp := network.NewVirtualNetworkGatewayListResultPage(network.VirtualNetworkGatewayListResult{Value: &[]network.VirtualNetworkGateway{vng}}, func(ctx context.Context, result network.VirtualNetworkGatewayListResult) (network.VirtualNetworkGatewayListResult, error) {
		return network.VirtualNetworkGatewayListResult{}, nil
	})
	vngcp := network.NewVirtualNetworkGatewayListConnectionsResultPage(network.VirtualNetworkGatewayListConnectionsResult{Value: &[]network.VirtualNetworkGatewayConnectionListEntity{vngc}}, func(ctx context.Context, result network.VirtualNetworkGatewayListConnectionsResult) (network.VirtualNetworkGatewayListConnectionsResult, error) {
		return network.VirtualNetworkGatewayListConnectionsResult{}, nil
	})

	n.EXPECT().ListAll(gomock.Any()).Return(vnp, nil)
	ngc.EXPECT().List(gomock.Any(), gomock.Any()).Return(vngp, nil)
	ngc.EXPECT().ListConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(vngcp, nil)
	return s
}

func TestNetworkVirtualNetworks(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkVirtualNetworks(), buildNetworkVirtualNetworksMock, client.TestOptions{})
}
