package wafregional

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafregionalClient(ctrl)

	tempRuleGroup := types.RuleGroup{}
	require.NoError(t, faker.FakeObject(&tempRuleGroup))

	tempRule := types.ActivatedRule{}
	require.NoError(t, faker.FakeObject(&tempRule))

	var tempTags []types.Tag
	require.NoError(t, faker.FakeObject(&tempTags))

	m.EXPECT().ListRuleGroups(
		gomock.Any(),
		&wafregional.ListRuleGroupsInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupSummary{{RuleGroupId: tempRuleGroup.RuleGroupId}},
		},
		nil,
	)

	m.EXPECT().GetRuleGroup(
		gomock.Any(),
		&wafregional.GetRuleGroupInput{RuleGroupId: tempRuleGroup.RuleGroupId},
		gomock.Any(),
	).Return(
		&wafregional.GetRuleGroupOutput{RuleGroup: &tempRuleGroup},
		nil,
	)

	m.EXPECT().ListActivatedRulesInRuleGroup(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&wafregional.ListActivatedRulesInRuleGroupOutput{
		ActivatedRules: []types.ActivatedRule{tempRule},
	}, nil)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:rulegroup/%v", *tempRuleGroup.RuleGroupId)),
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{
				TagList: tempTags,
			},
		},
		nil,
	)

	return client.Services{Wafregional: m}
}

func TestRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildRuleGroupsMock, client.TestOptions{})
}
