package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRegexPatternSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafV2Client(ctrl)

	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		var s types.RegexPatternSet
		if err := faker.FakeData(&s); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListRegexPatternSets(
			gomock.Any(),
			&wafv2.ListRegexPatternSetsInput{Scope: scope, Limit: aws.Int32(100)},
			gomock.Any(),
		).Return(
			&wafv2.ListRegexPatternSetsOutput{
				RegexPatternSets: []types.RegexPatternSetSummary{{Id: s.Id, Name: s.Name}},
			},
			nil,
		)

		m.EXPECT().GetRegexPatternSet(
			gomock.Any(),
			&wafv2.GetRegexPatternSetInput{Id: s.Id, Name: s.Name, Scope: scope},
			gomock.Any(),
		).Return(
			&wafv2.GetRegexPatternSetOutput{RegexPatternSet: &s},
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

	return client.Services{WafV2: m}
}

func TestWafV2RegexPatternSets(t *testing.T) {
	client.AwsMockTestHelper(t, RegexPatternSets(), buildRegexPatternSetsMock, client.TestOptions{})
}
