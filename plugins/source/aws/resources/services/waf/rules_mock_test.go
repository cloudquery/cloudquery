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

func buildWAFRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleSum := types.RuleSummary{}
	require.NoError(t, faker.FakeObject(&tempRuleSum))

	tempRule := types.Rule{}
	require.NoError(t, faker.FakeObject(&tempRule))

	var tempTags []types.Tag
	require.NoError(t, faker.FakeObject(&tempTags))

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
