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
				{
					SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
						NetworkSecurityGroup: &network.SecurityGroup{
							SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{},
						},
						RouteTable: &network.RouteTable{
							RouteTablePropertiesFormat: &network.RouteTablePropertiesFormat{},
						},
					},
				},
			},
		},
	}

	require.Nil(t, faker.FakeData(&vn.ID))
	require.Nil(t, faker.FakeData(&vn.Etag))
	require.Nil(t, faker.FakeData(&vn.Name))
	require.Nil(t, faker.FakeData(&vn.Tags))
	require.Nil(t, faker.FakeData(&vn.Type))
	require.Nil(t, faker.FakeData(&vn.Location))
	require.Nil(t, faker.FakeData(&vn.ExtendedLocation))
	require.Nil(t, faker.FakeData(&vn.ResourceGUID))
	require.Nil(t, faker.FakeData(&vn.ProvisioningState))
	require.Nil(t, faker.FakeData(&vn.AddressSpace))
	require.Nil(t, faker.FakeData(&vn.DhcpOptions))
	require.Nil(t, faker.FakeData(&vn.DdosProtectionPlan))
	require.Nil(t, faker.FakeData(&vn.EnableDdosProtection))
	require.Nil(t, faker.FakeData(&vn.EnableVMProtection))
	require.Nil(t, faker.FakeData(&vn.BgpCommunities))
	require.Nil(t, faker.FakeData(&vn.IPAllocations))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].ID))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].Name))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].Etag))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].AddressPrefix))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].AddressPrefixes))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NatGateway))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].ServiceEndpoints))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].IPAllocations))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].ResourceNavigationLinks))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].ServiceAssociationLinks))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].Delegations))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].Purpose))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].ProvisioningState))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].PrivateEndpointNetworkPolicies))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].PrivateLinkServiceNetworkPolicies))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].Etag))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].ID))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.ID))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.Name))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.Type))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.Tags))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.Etag))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.Location))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].NetworkSecurityGroup.ResourceGUID))

	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.ID))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.Type))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.Tags))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.Etag))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.Location))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.Name))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.DisableBgpRoutePropagation))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.ProvisioningState))
	require.Nil(t, faker.FakeData(&(*vn.Subnets)[0].RouteTable.ResourceGUID))
	require.Nil(t, faker.FakeData(&vn.VirtualNetworkPeerings))

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
