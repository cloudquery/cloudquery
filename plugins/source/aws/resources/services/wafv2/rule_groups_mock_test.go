package wafv2

import (
	"math/rand"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildWAFV2RuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafV2Client(ctrl)
	visibilityConfig := types.VisibilityConfig{}
	if err := faker.FakeObject(&visibilityConfig); err != nil {
		t.Fatal(err)
	}
	customRespBody := map[string]types.CustomResponseBody{}
	if err := faker.FakeObject(&customRespBody); err != nil {
		t.Fatal(err)
	}
	var labelSummaries []types.LabelSummary
	if err := faker.FakeObject(&labelSummaries); err != nil {
		t.Fatal(err)
	}
	overrideAction := types.OverrideAction{}
	if err := faker.FakeObject(&overrideAction); err != nil {
		t.Fatal(err)
	}
	action := types.RuleAction{}
	if err := faker.FakeObject(&action); err != nil {
		t.Fatal(err)
	}
	var labels []types.Label
	if err := faker.FakeObject(&labelSummaries); err != nil {
		t.Fatal(err)
	}
	rule := types.Rule{
		Name:             aws.String(faker.Name()),
		Priority:         rand.Int31(),
		Statement:        &types.Statement{AndStatement: &types.AndStatement{}},
		VisibilityConfig: &visibilityConfig,
		Action:           &action,
		OverrideAction:   &overrideAction,
		RuleLabels:       labels,
	}
	var tempPolicyOutput wafv2.GetPermissionPolicyOutput
	if err := faker.FakeObject(&tempPolicyOutput); err != nil {
		t.Fatal(err)
	}
	tempPolicyOutput.Policy = aws.String(`{"test": 1}`)
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempRuleGroupSum := types.RuleGroupSummary{}
		if err := faker.FakeObject(&tempRuleGroupSum); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListRuleGroups(gomock.Any(), &wafv2.ListRuleGroupsInput{Scope: scope}, gomock.Any()).Return(&wafv2.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
		}, nil)
		tempRuleGroup := types.RuleGroup{
			ARN:                  aws.String(faker.Word()),
			Capacity:             faker.RandomUnixTime(),
			Id:                   aws.String(faker.Word()),
			Name:                 aws.String(faker.Word()),
			VisibilityConfig:     &visibilityConfig,
			AvailableLabels:      labelSummaries,
			ConsumedLabels:       labelSummaries,
			CustomResponseBodies: customRespBody,
			Description:          aws.String(faker.Word()),
			LabelNamespace:       aws.String(faker.Word()),
			Rules:                []types.Rule{rule},
		}
		m.EXPECT().GetRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.GetRuleGroupOutput{
			RuleGroup: &tempRuleGroup,
		}, nil)
		m.EXPECT().GetPermissionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tempPolicyOutput, nil)
		m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
		}, nil)
	}

	return client.Services{WafV2: m}
}

func TestWafV2RuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildWAFV2RuleGroupsMock, client.TestOptions{})
}
