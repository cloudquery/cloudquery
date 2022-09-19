// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func createVirtualNetworkRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLVirtualNetworkRulesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			VirtualNetworkRules: mockClient,
		},
	}

	data := sql.VirtualNetworkRule{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewVirtualNetworkRuleListResultPage(sql.VirtualNetworkRuleListResult{Value: &[]sql.VirtualNetworkRule{data}}, func(ctx context.Context, result sql.VirtualNetworkRuleListResult) (sql.VirtualNetworkRuleListResult, error) {
		return sql.VirtualNetworkRuleListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
