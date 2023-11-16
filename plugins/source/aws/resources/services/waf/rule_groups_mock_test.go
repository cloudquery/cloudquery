package waf

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWAFRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleGroupSum := types.RuleGroupSummary{}
	require.NoError(t, faker.FakeObject(&tempRuleGroupSum))

	tempRuleGroup := types.RuleGroup{}
	require.NoError(t, faker.FakeObject(&tempRuleGroup))

	tempRule := types.ActivatedRule{}
	require.NoError(t, faker.FakeObject(&tempRule))

	var tempTags []types.Tag
	require.NoError(t, faker.FakeObject(&tempTags))

	m.EXPECT().ListRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListRuleGroupsOutput{
		RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
	}, nil)
	m.EXPECT().GetRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.GetRuleGroupOutput{
		RuleGroup: &tempRuleGroup,
	}, nil)
	m.EXPECT().ListActivatedRulesInRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListActivatedRulesInRuleGroupOutput{
		ActivatedRules: []types.ActivatedRule{tempRule},
	}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListTagsForResourceOutput{
		TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
	}, nil)

	return client.Services{Waf: m}
}

func TestWafRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildWAFRuleGroupsMock, client.TestOptions{})
}
