package route53resolver

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildFirewallRuleGroupAssociationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	frga := types.FirewallRuleGroupAssociation{}
	require.NoError(t, faker.FakeObject(&frga))

	m.EXPECT().ListFirewallRuleGroupAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListFirewallRuleGroupAssociationsOutput{
			FirewallRuleGroupAssociations: []types.FirewallRuleGroupAssociation{frga},
		}, nil)

	fdl := types.FirewallDomainList{}
	require.NoError(t, faker.FakeObject(&fdl))

	return client.Services{
		Route53resolver: m,
	}
}
func TestFirewallRuleGroupAssociations(t *testing.T) {
	client.AwsMockTestHelper(t, FirewallRuleGroupAssociations(), buildFirewallRuleGroupAssociationsMock, client.TestOptions{})
}
