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

func createManagedDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedDatabasesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedDatabases:                            mockClient,
			ManagedDatabaseVulnerabilityAssessments:     createManagedDatabaseVulnerabilityAssessmentsMock(t, ctrl).SQL.ManagedDatabaseVulnerabilityAssessments,
			ManagedDatabaseVulnerabilityAssessmentScans: createManagedDatabaseVulnerabilityAssessmentScansMock(t, ctrl).SQL.ManagedDatabaseVulnerabilityAssessmentScans,
		},
	}

	data := sql.ManagedDatabase{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewManagedDatabaseListResultPage(sql.ManagedDatabaseListResult{Value: &[]sql.ManagedDatabase{data}}, func(ctx context.Context, result sql.ManagedDatabaseListResult) (sql.ManagedDatabaseListResult, error) {
		return sql.ManagedDatabaseListResult{}, nil
	})

	mockClient.EXPECT().ListByInstance(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
