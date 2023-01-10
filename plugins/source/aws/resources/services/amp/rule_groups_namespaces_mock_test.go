package amp

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRuleGroupsNamespaces(t *testing.T, m *mocks.MockAmpClient) {
	var summary types.RuleGroupsNamespaceSummary
	if err := faker.FakeObject(&summary); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListRuleGroupsNamespaces(gomock.Any(), gomock.Any()).Return(
		&amp.ListRuleGroupsNamespacesOutput{
			RuleGroupsNamespaces: []types.RuleGroupsNamespaceSummary{summary},
		},
		nil,
	)

	var description types.RuleGroupsNamespaceDescription
	if err := faker.FakeObject(&description); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRuleGroupsNamespace(gomock.Any(), gomock.Any()).Return(
		&amp.DescribeRuleGroupsNamespaceOutput{
			RuleGroupsNamespace: &description,
		},
		nil,
	)
}
