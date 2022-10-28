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

func buildWAFSubscribedRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempSubscrRuleGroupSum := types.SubscribedRuleGroupSummary{}
	if err := faker.FakeObject(&tempSubscrRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSubscribedRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListSubscribedRuleGroupsOutput{
		RuleGroups: []types.SubscribedRuleGroupSummary{tempSubscrRuleGroupSum},
	}, nil)

	return client.Services{Waf: m}
}

func TestWafSubscribedRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubscribedRuleGroups(), buildWAFSubscribedRuleGroupsMock, client.TestOptions{})
}
