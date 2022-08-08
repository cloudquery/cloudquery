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

func buildRateBasedRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafRegionalClient(ctrl)

	var rule types.RateBasedRule
	if err := faker.FakeData(&rule); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRateBasedRules(
		gomock.Any(),
		&wafregional.ListRateBasedRulesInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListRateBasedRulesOutput{
			Rules: []types.RuleSummary{{RuleId: rule.RuleId}},
		},
		nil,
	)

	m.EXPECT().GetRateBasedRule(
		gomock.Any(),
		&wafregional.GetRateBasedRuleInput{RuleId: rule.RuleId},
		gomock.Any(),
	).Return(
		&wafregional.GetRateBasedRuleOutput{Rule: &rule},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:ratebasedrule/%v", *rule.RuleId)),
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{},
		nil,
	)

	return client.Services{WafRegional: m}
}

func TestRateBasedRules(t *testing.T) {
	client.AwsMockTestHelper(t, RateBasedRules(), buildRateBasedRulesMock, client.TestOptions{})
}
