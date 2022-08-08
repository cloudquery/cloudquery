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
			ServicePublicIPAddress: &network.PublicIPAddress{},
			LinkedPublicIPAddress:  &network.PublicIPAddress{},
		},
	}
	require.Nil(t, faker.FakeDataSkipFields(&pip, []string{"PublicIPAddressPropertiesFormat"}))
	require.Nil(t, faker.FakeDataSkipFields(pip.PublicIPAddressPropertiesFormat, []string{"PublicIPAllocationMethod", "PublicIPAddressVersion", "IPConfiguration",
		"ServicePublicIPAddress", "ProvisioningState", "MigrationPhase", "LinkedPublicIPAddress"}))
	require.Nil(t, faker.FakeDataSkipFields(pip.PublicIPAddressPropertiesFormat.IPConfiguration, []string{"IPConfigurationPropertiesFormat"}))
	require.Nil(t, faker.FakeDataSkipFields(pip.PublicIPAddressPropertiesFormat.ServicePublicIPAddress, []string{"PublicIPAddressPropertiesFormat"}))
	require.Nil(t, faker.FakeDataSkipFields(pip.PublicIPAddressPropertiesFormat.LinkedPublicIPAddress, []string{"PublicIPAddressPropertiesFormat"}))
	pip.PublicIPAddressPropertiesFormat.PublicIPAllocationMethod = "test"
	pip.PublicIPAddressPropertiesFormat.PublicIPAddressVersion = "test"
	pip.PublicIPAddressPropertiesFormat.ProvisioningState = "test"
	pip.PublicIPAddressPropertiesFormat.MigrationPhase = "test"
	fakeId := client.FakeResourceGroup + "/" + *pip.ID
	pip.ID = &fakeId
	ip := faker.IPv4()
	pip.IPAddress = &ip

	page := network.NewPublicIPAddressListResultPage(network.PublicIPAddressListResult{Value: &[]network.PublicIPAddress{pip}}, func(ctx context.Context, result network.PublicIPAddressListResult) (network.PublicIPAddressListResult, error) {
		return network.PublicIPAddressListResult{}, nil
	})
	pips.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestPublicIpAddresses(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkPublicIPAddresses(), buildNetworkPublicIpAddressesMock, client.TestOptions{})
}
