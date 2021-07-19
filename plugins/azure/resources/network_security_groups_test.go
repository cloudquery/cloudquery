package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNetworkSecurityGroupsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	n := mocks.NewMockSecurityGroupsClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			SecurityGroups: n,
		},
	}

	sg := network.SecurityGroup{SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
		Subnets: &[]network.Subnet{
			{
				SubnetPropertiesFormat: &network.SubnetPropertiesFormat{},
			},
		},
		NetworkInterfaces: &[]network.Interface{},
	}}
	require.Nil(t, faker.FakeData(&sg.ID))
	require.Nil(t, faker.FakeData(&sg.Etag))
	require.Nil(t, faker.FakeData(&sg.Name))
	require.Nil(t, faker.FakeData(&sg.Tags))
	require.Nil(t, faker.FakeData(&sg.Type))
	require.Nil(t, faker.FakeData(&sg.Location))
	require.Nil(t, faker.FakeData(&sg.SecurityRules))
	require.Nil(t, faker.FakeData(&sg.DefaultSecurityRules))
	require.Nil(t, faker.FakeData(&sg.FlowLogs))
	require.Nil(t, faker.FakeData(&sg.ResourceGUID))
	require.Nil(t, faker.FakeData(&sg.ProvisioningState))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].ID))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].Name))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].Etag))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].AddressPrefix))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].AddressPrefixes))
	//require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].NetworkSecurityGroup.ID))
	//require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].RouteTable))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].NatGateway))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].ServiceEndpoints))
	//require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].ServiceEndpointPolicies))
	//require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].PrivateEndpoints))
	//require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].IPConfigurations))
	//require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].IPConfigurationProfiles))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].IPAllocations))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].ResourceNavigationLinks))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].ServiceAssociationLinks))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].Delegations))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].Purpose))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].ProvisioningState))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].PrivateEndpointNetworkPolicies))
	require.Nil(t, faker.FakeData(&(*sg.Subnets)[0].PrivateLinkServiceNetworkPolicies))

	fakeId := fakeResourceGroup + "/" + *sg.ID
	sg.ID = &fakeId

	page := network.NewSecurityGroupListResultPage(network.SecurityGroupListResult{Value: &[]network.SecurityGroup{sg}}, func(ctx context.Context, result network.SecurityGroupListResult) (network.SecurityGroupListResult, error) {
		return network.SecurityGroupListResult{}, nil
	})
	n.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkSecurityGroups(t *testing.T) {
	azureTestHelper(t, resources.NetworkSecurityGroups(), buildNetworkSecurityGroupsMock)
}
