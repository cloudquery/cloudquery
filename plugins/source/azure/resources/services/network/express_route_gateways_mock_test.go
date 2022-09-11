// Auto generated code - DO NOT EDIT.

package network

import (
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

func TestNetworkExpressRouteGateways(t *testing.T) {
	client.MockTestHelper(t, ExpressRouteGateways(), createExpressRouteGatewaysMock)
}

func createExpressRouteGatewaysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRouteGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRouteGateways: mockClient,
		},
	}

	data := network.ExpressRouteGateway{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := network.ExpressRouteGatewayList{Value: &[]network.ExpressRouteGateway{data}}

	mockClient.EXPECT().ListBySubscription(gomock.Any()).Return(result, nil)
	return s
}
