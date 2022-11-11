package waf

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildWAFRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleGroupSum := types.RuleGroupSummary{}
	if err := faker.FakeObject(&tempRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	tempRuleGroup := types.RuleGroup{}
	if err := faker.FakeObject(&tempRuleGroup); err != nil {
		t.Fatal(err)
	}
	tempRule := types.ActivatedRule{}
	if err := faker.FakeObject(&tempRule); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
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
