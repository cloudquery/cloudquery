// Auto generated code - DO NOT EDIT.

package network

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func TestNetworkPublicIPAddresses(t *testing.T) {
	client.AzureMockTestHelper(t, PublicIPAddresses(), createPublicIPAddressesMock, client.TestOptions{})
}

func createPublicIPAddressesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkPublicIPAddressesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			PublicIPAddresses: mockClient,
		},
	}

	data := network.PublicIPAddress{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := network.NewPublicIPAddressListResultPage(network.PublicIPAddressListResult{Value: &[]network.PublicIPAddress{data}}, func(ctx context.Context, result network.PublicIPAddressListResult) (network.PublicIPAddressListResult, error) {
		return network.PublicIPAddressListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
