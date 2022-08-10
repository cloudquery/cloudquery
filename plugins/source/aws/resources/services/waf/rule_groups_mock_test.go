package waf

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWAFRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleGroupSum := types.RuleGroupSummary{}
	if err := faker.FakeData(&tempRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	tempRuleGroup := types.RuleGroup{}
	if err := faker.FakeData(&tempRuleGroup); err != nil {
		t.Fatal(err)
	}
	tempRule := types.ActivatedRule{}
	if err := faker.FakeData(&tempRule); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeData(&tempTags); err != nil {
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
	client.AwsMockTestHelper(t, WafRuleGroups(), buildWAFRuleGroupsMock, client.TestOptions{})
}
