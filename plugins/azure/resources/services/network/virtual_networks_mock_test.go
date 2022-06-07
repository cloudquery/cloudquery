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

func buildNetworkVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	n := mocks.NewMockVirtualNetworksClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			VirtualNetworks: n,
		},
	}

	vn := fakeVirtualNetwork(t)
	page := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{vn}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})
	n.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkVirtualNetworks(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkVirtualNetworks(), buildNetworkVirtualNetworksMock, client.TestOptions{})
}
