// +build mock

package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"

	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWAFV2ManagedRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafV2Client(ctrl)
	tempManagedRuleGroupSum := types.ManagedRuleGroupSummary{}
	if err := faker.FakeData(&tempManagedRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	tempDescribeManagedRuleGroup := wafv2.DescribeManagedRuleGroupOutput{}
	if err := faker.FakeData(&tempDescribeManagedRuleGroup); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListAvailableManagedRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListAvailableManagedRuleGroupsOutput{
		ManagedRuleGroups: []types.ManagedRuleGroupSummary{tempManagedRuleGroupSum},
	}, nil)
	m.EXPECT().DescribeManagedRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tempDescribeManagedRuleGroup, nil)

	return client.Services{WafV2: m}
}

func TestWafV2ManagedRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Wafv2ManagedRuleGroups(), buildWAFV2ManagedRuleGroupsMock, client.TestOptions{})
}
