package networkfirewall

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildFirewallPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	fpm := types.FirewallPolicyMetadata{}
	err := faker.FakeObject(&fpm)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListFirewallPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListFirewallPoliciesOutput{
			FirewallPolicies: []types.FirewallPolicyMetadata{fpm},
		}, nil)

	fp := networkfirewall.DescribeFirewallPolicyOutput{}
	if err := faker.FakeObject(&fp); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeFirewallPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&fp, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestFirewallPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, FirewallPolicies(), buildFirewallPoliciesMock, client.TestOptions{})
}
