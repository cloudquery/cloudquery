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

func buildWAFV2RuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafv2Client(ctrl)
	visibilityConfig := types.VisibilityConfig{}
	require.NoError(t, faker.FakeObject(&visibilityConfig))

	customRespBody := map[string]types.CustomResponseBody{}
	require.NoError(t, faker.FakeObject(&customRespBody))

	var labelSummaries []types.LabelSummary
	require.NoError(t, faker.FakeObject(&labelSummaries))

	overrideAction := types.OverrideAction{}
	require.NoError(t, faker.FakeObject(&overrideAction))

	action := types.RuleAction{}
	require.NoError(t, faker.FakeObject(&action))

	var labels []types.Label
	require.NoError(t, faker.FakeObject(&labelSummaries))

	rule := types.Rule{}
	require.NoError(t, faker.FakeObject(&rule))

	rule.VisibilityConfig = &visibilityConfig
	rule.Action = &action
	rule.OverrideAction = &overrideAction
	rule.RuleLabels = labels
	var tempPolicyOutput wafv2.GetPermissionPolicyOutput
	require.NoError(t, faker.FakeObject(&tempPolicyOutput))

	tempPolicyOutput.Policy = aws.String(`{"test": 1}`)
	var tempTags []types.Tag
	require.NoError(t, faker.FakeObject(&tempTags))

	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempRuleGroupSum := types.RuleGroupSummary{}
		require.NoError(t, faker.FakeObject(&tempRuleGroupSum))
		m.EXPECT().ListRuleGroups(gomock.Any(), &wafv2.ListRuleGroupsInput{Scope: scope}, gomock.Any()).Return(&wafv2.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
		}, nil)

		tempRuleGroup := types.RuleGroup{}
		require.NoError(t, faker.FakeObject(&tempRuleGroup))
		m.EXPECT().GetRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.GetRuleGroupOutput{
			RuleGroup: &tempRuleGroup,
		}, nil)
		m.EXPECT().GetPermissionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tempPolicyOutput, nil)
		m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
		}, nil)
	}

	return client.Services{Wafv2: m}
}

func TestWafV2RuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildWAFV2RuleGroupsMock, client.TestOptions{})
}
