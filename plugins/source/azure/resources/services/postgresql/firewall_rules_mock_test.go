// Auto generated code - DO NOT EDIT.

package postgresql

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
)

func createFirewallRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockPostgreSQLFirewallRulesClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQLClient{
			FirewallRules: mockClient,
		},
	}

	data := postgresql.FirewallRule{}
	require.Nil(t, faker.FakeObject(&data))

	result := postgresql.FirewallRuleListResult{Value: &[]postgresql.FirewallRule{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
