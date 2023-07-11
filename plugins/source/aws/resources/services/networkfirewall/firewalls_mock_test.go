package networkfirewall

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildFirewallsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	fm := types.FirewallMetadata{}
	require.NoError(t, faker.FakeObject(&fm))

	m.EXPECT().ListFirewalls(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListFirewallsOutput{
			Firewalls: []types.FirewallMetadata{fm},
		}, nil)

	fo := networkfirewall.DescribeFirewallOutput{}
	require.NoError(t, faker.FakeObject(&fo))

	m.EXPECT().DescribeFirewall(gomock.Any(), gomock.Any(), gomock.Any()).Return(&fo, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestFirewalls(t *testing.T) {
	client.AwsMockTestHelper(t, Firewalls(), buildFirewallsMock, client.TestOptions{})
}
