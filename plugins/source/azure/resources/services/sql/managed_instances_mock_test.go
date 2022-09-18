// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func TestSQLManagedInstances(t *testing.T) {
	client.MockTestHelper(t, ManagedInstances(), createManagedInstancesMock)
}

func createManagedInstancesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedInstancesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedInstances:                            mockClient,
			ManagedDatabases:                            createManagedDatabasesMock(t, ctrl).SQL.ManagedDatabases,
			ManagedDatabaseVulnerabilityAssessments:     createManagedDatabaseVulnerabilityAssessmentsMock(t, ctrl).SQL.ManagedDatabaseVulnerabilityAssessments,
			ManagedDatabaseVulnerabilityAssessmentScans: createManagedDatabaseVulnerabilityAssessmentScansMock(t, ctrl).SQL.ManagedDatabaseVulnerabilityAssessmentScans,
			ManagedInstanceVulnerabilityAssessments:     createManagedInstanceVulnerabilityAssessmentsMock(t, ctrl).SQL.ManagedInstanceVulnerabilityAssessments,
			ManagedInstanceEncryptionProtectors:         createManagedInstanceEncryptionProtectorsMock(t, ctrl).SQL.ManagedInstanceEncryptionProtectors,
		},
	}

	data := sql.ManagedInstance{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewManagedInstanceListResultPage(sql.ManagedInstanceListResult{Value: &[]sql.ManagedInstance{data}}, func(ctx context.Context, result sql.ManagedInstanceListResult) (sql.ManagedInstanceListResult, error) {
		return sql.ManagedInstanceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
