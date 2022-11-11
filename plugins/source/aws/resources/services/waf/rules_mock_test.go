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

func buildWAFRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleSum := types.RuleSummary{}
	if err := faker.FakeObject(&tempRuleSum); err != nil {
		t.Fatal(err)
	}
	tempRule := types.Rule{}
	if err := faker.FakeObject(&tempRule); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListRulesOutput{
		Rules: []types.RuleSummary{tempRuleSum},
	}, nil)
	m.EXPECT().GetRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.GetRuleOutput{
		Rule: &tempRule,
	}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListTagsForResourceOutput{
		TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
	}, nil)

	return client.Services{Waf: m}
}

func TestWafRules(t *testing.T) {
	client.AwsMockTestHelper(t, Rules(), buildWAFRulesMock, client.TestOptions{})
}
