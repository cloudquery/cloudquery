//go:generate mockgen -destination=./mocks/network.go -package=mocks . VirtualNetworksClient,SecurityGroupsClient,WatchersClient,PublicIPAddressesClient,InterfacesClient
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

func fakeInterfaceIPConfiguration(t *testing.T) network.InterfaceIPConfiguration {
	subnet := fakeSubnet(t)
	sb := network.InterfaceIPConfiguration{
		InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
			PrivateIPAllocationMethod: network.IPAllocationMethodDynamic,
			Subnet:                    &subnet,
		},
	}
	var bap network.BackendAddressPool
	require.NoError(t, faker.FakeDataSkipFields(&bap, []string{"BackendAddressPoolPropertiesFormat"}))
	sb.InterfaceIPConfigurationPropertiesFormat.LoadBalancerBackendAddressPools = &[]network.BackendAddressPool{bap}

	var aap network.ApplicationGatewayBackendAddressPool
	require.NoError(t, faker.FakeDataSkipFields(&aap, []string{"ApplicationGatewayBackendAddressPoolPropertiesFormat"}))
	sb.InterfaceIPConfigurationPropertiesFormat.ApplicationGatewayBackendAddressPools = &[]network.ApplicationGatewayBackendAddressPool{aap}

	var vnt network.VirtualNetworkTap
	require.NoError(t, faker.FakeDataSkipFields(&vnt, []string{"VirtualNetworkTapPropertiesFormat"}))
	sb.InterfaceIPConfigurationPropertiesFormat.VirtualNetworkTaps = &[]network.VirtualNetworkTap{vnt}

	var inr network.InboundNatRule
	require.NoError(t, faker.FakeDataSkipFields(&inr, []string{"InboundNatRulePropertiesFormat"}))
	sb.InterfaceIPConfigurationPropertiesFormat.LoadBalancerInboundNatRules = &[]network.InboundNatRule{inr}

	var pip network.PublicIPAddress
	require.NoError(t, faker.FakeDataSkipFields(&pip, []string{"PublicIPAddressPropertiesFormat"}))
	sb.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress = &pip

	require.NoError(t, faker.FakeDataSkipFields(&sb, []string{"ProvisioningState", "InterfaceIPConfigurationPropertiesFormat"}))
	require.NoError(t, faker.FakeDataSkipFields(sb.InterfaceIPConfigurationPropertiesFormat, []string{
		"ApplicationGatewayBackendAddressPools",
		"LoadBalancerBackendAddressPools",
		"LoadBalancerInboundNatRules",
		"PrivateIPAddressVersion",
		"PrivateIPAllocationMethod",
		"ProvisioningState",
		"PublicIPAddress",
		"Subnet",
		"VirtualNetworkTaps",
	}))
	sb.InterfaceIPConfigurationPropertiesFormat.ProvisioningState = network.ProvisioningStateSucceeded
	sb.InterfaceIPConfigurationPropertiesFormat.PrivateIPAddressVersion = network.IPVersionIPv4
	var asg network.ApplicationSecurityGroup
	require.NoError(t, faker.FakeDataSkipFields(&asg, []string{"ApplicationSecurityGroupPropertiesFormat"}))
	sb.InterfaceIPConfigurationPropertiesFormat.ApplicationSecurityGroups = &[]network.ApplicationSecurityGroup{asg}
	return sb
}

func buildNetworkInterfacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	ic := mocks.NewMockInterfacesClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			Interfaces: ic,
		},
	}

	i := network.Interface{
		InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
			NetworkSecurityGroup: &network.SecurityGroup{},
			PrivateEndpoint:      &network.PrivateEndpoint{},
			TapConfigurations:    &[]network.InterfaceTapConfiguration{},
			PrivateLinkService:   &network.PrivateLinkService{},
			IPConfigurations: &[]network.InterfaceIPConfiguration{
				fakeInterfaceIPConfiguration(t),
			},
		},
	}

	require.Nil(t, faker.FakeData(&i.ID))
	require.Nil(t, faker.FakeData(&i.Etag))
	require.Nil(t, faker.FakeData(&i.ExtendedLocation))
	require.Nil(t, faker.FakeData(&i.Location))
	require.Nil(t, faker.FakeData(&i.Name))
	require.Nil(t, faker.FakeData(&i.Tags))
	require.Nil(t, faker.FakeData(&i.Type))
	require.Nil(t, faker.FakeData(&i.PrivateEndpoint.ID))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.DNSSettings))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.DscpConfiguration))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.HostedWorkloads))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.ResourceGUID))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.MacAddress))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.VirtualMachine))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.Primary))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.EnableAcceleratedNetworking))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.EnableIPForwarding))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.MigrationPhase))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.NicType))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.NetworkSecurityGroup.ID))
	fakeId := client.FakeResourceGroup + "/" + *i.ID
	i.ID = &fakeId

	page := network.NewInterfaceListResultPage(network.InterfaceListResult{Value: &[]network.Interface{i}}, func(ctx context.Context, result network.InterfaceListResult) (network.InterfaceListResult, error) {
		return network.InterfaceListResult{}, nil
	})
	ic.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkInterfaces(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkInterfaces(), buildNetworkInterfacesMock, client.TestOptions{})
}
