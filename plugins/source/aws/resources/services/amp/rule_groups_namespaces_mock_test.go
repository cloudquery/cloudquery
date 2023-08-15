package amp

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRuleGroupsNamespaces(t *testing.T, m *mocks.MockAmpClient) {
	var summary types.RuleGroupsNamespaceSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListRuleGroupsNamespaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amp.ListRuleGroupsNamespacesOutput{
			RuleGroupsNamespaces: []types.RuleGroupsNamespaceSummary{summary},
		},
		nil,
	)

	var description types.RuleGroupsNamespaceDescription
	require.NoError(t, faker.FakeObject(&description))

	m.EXPECT().DescribeRuleGroupsNamespace(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amp.DescribeRuleGroupsNamespaceOutput{
			RuleGroupsNamespace: &description,
		},
		nil,
	)
}
