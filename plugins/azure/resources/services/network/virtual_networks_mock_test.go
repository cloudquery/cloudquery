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
)

func fakeSubnet(t *testing.T) network.Subnet {
	sb := network.Subnet{
		SubnetPropertiesFormat: &network.SubnetPropertiesFormat{},
	}
	if err := faker.FakeDataSkipFields(&sb, []string{"ProvisioningState", "SubnetPropertiesFormat"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(sb.SubnetPropertiesFormat, []string{"ApplicationGatewayIPConfigurations",
		"RouteTable",
		"NetworkSecurityGroup",
		"ServiceEndpointPolicies",
		"PrivateEndpoints",
		"IPConfigurations",
		"IPConfigurationProfiles",
		"ProvisioningState",
		"PrivateEndpointNetworkPolicies",
		"PrivateLinkServiceNetworkPolicies"}); err != nil {
		t.Fatal(err)
	}
	return sb
}

func buildNetworkVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	n := mocks.NewMockVirtualNetworksClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			VirtualNetworks: n,
		},
	}

	vn := network.VirtualNetwork{
		VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
			Subnets: &[]network.Subnet{
				fakeSubnet(t),
			},
		},
	}
	if err := faker.FakeDataSkipFields(&vn, []string{"VirtualNetworkPropertiesFormat"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(vn.VirtualNetworkPropertiesFormat, []string{"Subnets", "ProvisioningState"}); err != nil {
		t.Fatal(err)
	}

	fakeId := client.FakeResourceGroup + "/" + *vn.ID
	vn.ID = &fakeId
	vn.DhcpOptions.DNSServers = &[]string{faker.IPv4()}

	page := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{vn}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})
	n.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkVirtualNetworks(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkVirtualNetworks(), buildNetworkVirtualNetworksMock, client.TestOptions{})
}
