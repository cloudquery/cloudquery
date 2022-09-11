// Auto generated code - DO NOT EDIT.

package sql

import (
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

	result := data

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
