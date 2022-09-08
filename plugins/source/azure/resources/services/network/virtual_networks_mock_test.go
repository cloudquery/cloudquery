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

func TestNetworkVirtualNetworks(t *testing.T) {
	client.MockTestHelper(t, VirtualNetworks(), createVirtualNetworksMock)
}

func createVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworksClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworks: mockClient,
		},
	}

	data := network.VirtualNetwork{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{data}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
