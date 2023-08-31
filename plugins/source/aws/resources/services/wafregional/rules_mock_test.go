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

func buildRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafregionalClient(ctrl)

	var r types.Rule
	require.NoError(t, faker.FakeObject(&r))

	m.EXPECT().ListRules(
		gomock.Any(),
		&wafregional.ListRulesInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListRulesOutput{
			Rules: []types.RuleSummary{{RuleId: r.RuleId}},
		},
		nil,
	)

	m.EXPECT().GetRule(
		gomock.Any(),
		&wafregional.GetRuleInput{RuleId: r.RuleId},
		gomock.Any(),
	).Return(
		&wafregional.GetRuleOutput{Rule: &r},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:rule/%v", *r.RuleId)),
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{},
		nil,
	)

	return client.Services{Wafregional: m}
}

func TestRules(t *testing.T) {
	client.AwsMockTestHelper(t, Rules(), buildRulesMock, client.TestOptions{})
}
