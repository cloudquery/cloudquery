package wafregional

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafRegionalClient(ctrl)

	var g types.RuleGroup
	if err := faker.FakeData(&g); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRuleGroups(
		gomock.Any(),
		&wafregional.ListRuleGroupsInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupSummary{{RuleGroupId: g.RuleGroupId}},
		},
		nil,
	)

	m.EXPECT().GetRuleGroup(
		gomock.Any(),
		&wafregional.GetRuleGroupInput{RuleGroupId: g.RuleGroupId},
		gomock.Any(),
	).Return(
		&wafregional.GetRuleGroupOutput{RuleGroup: &g},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:rulegroup/%v", *g.RuleGroupId)),
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{},
		nil,
	)

	return client.Services{WafRegional: m}
}

func TestRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildRuleGroupsMock, client.TestOptions{})
}
