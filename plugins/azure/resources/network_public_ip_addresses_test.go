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

func buildNetworkPublicIpAddressesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	pips := mocks.NewMockPublicIPAddressesClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			PublicIPAddresses: pips,
		},
	}

	pip := network.PublicIPAddress{
		PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
			IPConfiguration: &network.IPConfiguration{
				IPConfigurationPropertiesFormat: &network.IPConfigurationPropertiesFormat{
					Subnet:          &network.Subnet{},
					PublicIPAddress: &network.PublicIPAddress{},
				},
			},
		},
	}
	require.Nil(t, faker.FakeData(&pip.ExtendedLocation))
	require.Nil(t, faker.FakeData(&pip.Sku))
	require.Nil(t, faker.FakeData(&pip.Etag))
	require.Nil(t, faker.FakeData(&pip.Zones))
	require.Nil(t, faker.FakeData(&pip.ID))
	require.Nil(t, faker.FakeData(&pip.Name))
	require.Nil(t, faker.FakeData(&pip.Type))
	require.Nil(t, faker.FakeData(&pip.Location))
	require.Nil(t, faker.FakeData(&pip.Tags))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.PublicIPAllocationMethod))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.PublicIPAddressVersion))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.DNSSettings))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.DdosSettings))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPTags))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPAddress))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.PublicIPPrefix))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IdleTimeoutInMinutes))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.ResourceGUID))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.ProvisioningState))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.ID))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.Etag))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.Name))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PrivateIPAddress))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PrivateIPAllocationMethod))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.ProvisioningState))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.Subnet.ID))
	require.Nil(t, faker.FakeData(&pip.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress.ID))

	fakeId := fakeResourceGroup + "/" + *pip.ID
	pip.ID = &fakeId

	page := network.NewPublicIPAddressListResultPage(network.PublicIPAddressListResult{Value: &[]network.PublicIPAddress{pip}}, func(ctx context.Context, result network.PublicIPAddressListResult) (network.PublicIPAddressListResult, error) {
		return network.PublicIPAddressListResult{}, nil
	})
	pips.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestPublicIpAddresses(t *testing.T) {
	azureTestHelper(t, resources.NetworkPublicIPAddresses(), buildNetworkPublicIpAddressesMock)
}
