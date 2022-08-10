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

func buildNetworkRouteFiltersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	rfc := mocks.NewMockRouteFiltersClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			RouteFilters: rfc,
		},
	}

	tid := "test"
	rfr := network.RouteFilterRule{ID: &tid}
	require.Nil(t, faker.FakeData(&rfr.Etag))
	require.Nil(t, faker.FakeData(&rfr.Location))
	require.Nil(t, faker.FakeData(&rfr.Name))
	require.Nil(t, faker.FakeData(&rfr.RouteFilterRulePropertiesFormat))

	rf := network.RouteFilter{ID: &tid}
	require.Nil(t, faker.FakeData(&rf.Etag))
	require.Nil(t, faker.FakeData(&rf.Location))
	require.Nil(t, faker.FakeData(&rf.Name))
	require.Nil(t, faker.FakeData(&rf.Tags))
	require.Nil(t, faker.FakeData(&rf.Type))
	require.Nil(t, faker.FakeData(&rf.RouteFilterPropertiesFormat))
	rf.RouteFilterPropertiesFormat.Rules = &[]network.RouteFilterRule{rfr}
	fakeId := client.FakeResourceGroup + "/" + *rf.ID
	rf.ID = &fakeId

	page := network.NewRouteFilterListResultPage(network.RouteFilterListResult{Value: &[]network.RouteFilter{rf}}, func(ctx context.Context, result network.RouteFilterListResult) (network.RouteFilterListResult, error) {
		return network.RouteFilterListResult{}, nil
	})
	rfc.EXPECT().List(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkRouteFilters(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkRouteFilters(), buildNetworkRouteFiltersMock, client.TestOptions{})
}
