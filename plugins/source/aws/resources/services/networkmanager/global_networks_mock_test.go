package networkmanager

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNetworksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkmanagerClient(ctrl)

	network := types.GlobalNetwork{}
	require.NoError(t, faker.FakeObject(&network))

	m.EXPECT().DescribeGlobalNetworks(gomock.Any(), &networkmanager.DescribeGlobalNetworksInput{}, gomock.Any()).Return(
		&networkmanager.DescribeGlobalNetworksOutput{GlobalNetworks: []types.GlobalNetwork{network}}, nil)

	site := types.Site{}
	require.NoError(t, faker.FakeObject(&site))

	m.EXPECT().GetSites(gomock.Any(), &networkmanager.GetSitesInput{
		GlobalNetworkId: network.GlobalNetworkId,
	}, gomock.Any()).Return(
		&networkmanager.GetSitesOutput{Sites: []types.Site{site}}, nil)

	link := types.Link{}
	require.NoError(t, faker.FakeObject(&link))

	m.EXPECT().GetLinks(gomock.Any(), &networkmanager.GetLinksInput{
		GlobalNetworkId: network.GlobalNetworkId,
	}, gomock.Any()).Return(
		&networkmanager.GetLinksOutput{Links: []types.Link{link}}, nil)

	registration := types.TransitGatewayRegistration{}
	require.NoError(t, faker.FakeObject(&registration))

	m.EXPECT().GetTransitGatewayRegistrations(gomock.Any(), &networkmanager.GetTransitGatewayRegistrationsInput{
		GlobalNetworkId: network.GlobalNetworkId,
	}, gomock.Any()).Return(
		&networkmanager.GetTransitGatewayRegistrationsOutput{TransitGatewayRegistrations: []types.TransitGatewayRegistration{registration}}, nil)

	return client.Services{Networkmanager: m}
}

func TestGlobalNetworks(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalNetworks(), buildNetworksMock, client.TestOptions{
		Region: "us-west-2",
	})
}
