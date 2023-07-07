package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIpsetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafv2Client(ctrl)

	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		var s types.IPSet
		require.NoError(t, faker.FakeObject(&s))
		s.Addresses = []string{"192.0.2.0/24", "2001:db8::/32"}
		m.EXPECT().ListIPSets(
			gomock.Any(),
			&wafv2.ListIPSetsInput{Scope: scope, Limit: aws.Int32(100)},
			gomock.Any(),
		).Return(
			&wafv2.ListIPSetsOutput{
				IPSets: []types.IPSetSummary{{Id: s.Id, Name: s.Name}},
			},
			nil,
		)

		m.EXPECT().GetIPSet(
			gomock.Any(),
			&wafv2.GetIPSetInput{Name: s.Name, Id: s.Id, Scope: scope},
			gomock.Any(),
		).Return(
			&wafv2.GetIPSetOutput{IPSet: &s},
			nil,
		)

		m.EXPECT().ListTagsForResource(
			gomock.Any(),
			&wafv2.ListTagsForResourceInput{ResourceARN: s.ARN},
			gomock.Any(),
		).Return(
			&wafv2.ListTagsForResourceOutput{
				TagInfoForResource: &types.TagInfoForResource{
					TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
				},
			},
			nil,
		)
	}

	return client.Services{Wafv2: m}
}

func TestWafV2IPSets(t *testing.T) {
	client.AwsMockTestHelper(t, Ipsets(), buildIpsetsMock, client.TestOptions{})
}
