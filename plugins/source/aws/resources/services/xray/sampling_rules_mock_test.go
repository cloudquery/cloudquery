package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSamplingRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	test := "test"

	var rule types.SamplingRuleRecord
	if err := faker.FakeObject(&rule); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetSamplingRules(
		gomock.Any(),
		&xray.GetSamplingRulesInput{},
		gomock.Any(),
	).Return(
		&xray.GetSamplingRulesOutput{
			SamplingRuleRecords: []types.SamplingRuleRecord{
				rule,
			},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&xray.ListTagsForResourceInput{ResourceARN: rule.SamplingRule.RuleARN},
		gomock.Any(),
	).Return(
		&xray.ListTagsForResourceOutput{
			Tags: []types.Tag{
				{
					Key:   &test,
					Value: &test,
				},
			},
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestXraySamplingRules(t *testing.T) {
	client.AwsMockTestHelper(t, SamplingRules(), buildSamplingRules, client.TestOptions{})
}
