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

func TestNetworkRouteFilters(t *testing.T) {
	client.MockTestHelper(t, RouteFilters(), createRouteFiltersMock)
}

func createRouteFiltersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkRouteFiltersClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			RouteFilters: mockClient,
		},
	}

	data := network.RouteFilter{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := network.NewRouteFilterListResultPage(network.RouteFilterListResult{Value: &[]network.RouteFilter{data}}, func(ctx context.Context, result network.RouteFilterListResult) (network.RouteFilterListResult, error) {
		return network.RouteFilterListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
