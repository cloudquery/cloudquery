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

func TestSQLServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), createServersMock)
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServersClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Servers:                              mockClient,
			Databases:                            createDatabasesMock(t, ctrl).SQL.Databases,
			DatabaseBlobAuditingPolicies:         createDatabaseBlobAuditingPoliciesMock(t, ctrl).SQL.DatabaseBlobAuditingPolicies,
			DatabaseVulnerabilityAssessments:     createDatabaseVulnerabilityAssessmentsMock(t, ctrl).SQL.DatabaseVulnerabilityAssessments,
			DatabaseVulnerabilityAssessmentScans: createDatabaseVulnerabilityAssessmentScansMock(t, ctrl).SQL.DatabaseVulnerabilityAssessmentScans,
			BackupLongTermRetentionPolicies:      createBackupLongTermRetentionPoliciesMock(t, ctrl).SQL.BackupLongTermRetentionPolicies,
			DatabaseThreatDetectionPolicies:      createDatabaseThreatDetectionPoliciesMock(t, ctrl).SQL.DatabaseThreatDetectionPolicies,
			TransparentDataEncryptions:           createTransparentDataEncryptionsMock(t, ctrl).SQL.TransparentDataEncryptions,
			EncryptionProtectors:                 createEncryptionProtectorsMock(t, ctrl).SQL.EncryptionProtectors,
			VirtualNetworkRules:                  createVirtualNetworkRulesMock(t, ctrl).SQL.VirtualNetworkRules,
			FirewallRules:                        createFirewallRulesMock(t, ctrl).SQL.FirewallRules,
			ServerAdmins:                         createServerAdminsMock(t, ctrl).SQL.ServerAdmins,
			ServerBlobAuditingPolicies:           createServerBlobAuditingPoliciesMock(t, ctrl).SQL.ServerBlobAuditingPolicies,
			ServerDevOpsAuditingSettings:         createServerDevOpsAuditingSettingsMock(t, ctrl).SQL.ServerDevOpsAuditingSettings,
			ServerVulnerabilityAssessments:       createServerVulnerabilityAssessmentsMock(t, ctrl).SQL.ServerVulnerabilityAssessments,
			ServerSecurityAlertPolicies:          createServerSecurityAlertPoliciesMock(t, ctrl).SQL.ServerSecurityAlertPolicies,
		},
	}

	data := sql.Server{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewServerListResultPage(sql.ServerListResult{Value: &[]sql.Server{data}}, func(ctx context.Context, result sql.ServerListResult) (sql.ServerListResult, error) {
		return sql.ServerListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
