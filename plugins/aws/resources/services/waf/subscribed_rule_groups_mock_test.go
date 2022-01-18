//go:build mock
// +build mock

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

func buildWAFSubscribedRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempSubscrRuleGroupSum := types.SubscribedRuleGroupSummary{}
	if err := faker.FakeData(&tempSubscrRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSubscribedRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListSubscribedRuleGroupsOutput{
		RuleGroups: []types.SubscribedRuleGroupSummary{tempSubscrRuleGroupSum},
	}, nil)

	return client.Services{Waf: m}
}

func TestWafSubscribedRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, WafSubscribedRuleGroups(), buildWAFSubscribedRuleGroupsMock, client.TestOptions{})
}
