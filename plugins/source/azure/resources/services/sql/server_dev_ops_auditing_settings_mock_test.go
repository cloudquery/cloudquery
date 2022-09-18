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

func createServerDevOpsAuditingSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerDevOpsAuditingSettingsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerDevOpsAuditingSettings: mockClient,
		},
	}

	data := sql.ServerDevOpsAuditingSettings{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewServerDevOpsAuditingSettingsListResultPage(sql.ServerDevOpsAuditingSettingsListResult{Value: &[]sql.ServerDevOpsAuditingSettings{data}}, func(ctx context.Context, result sql.ServerDevOpsAuditingSettingsListResult) (sql.ServerDevOpsAuditingSettingsListResult, error) {
		return sql.ServerDevOpsAuditingSettingsListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
