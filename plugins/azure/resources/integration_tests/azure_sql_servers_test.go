package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSQLServers(t *testing.T) {
	table := resources.SQLServers()

	awsTestIntegrationHelper(t, table, []string{"azure_sql_servers.tf", "azure_storage_accounts.tf", "networks.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                fmt.Sprintf("mssql-1-%s-%s", res.Prefix, res.Suffix),
					"version":             "12.0",
					"administrator_login": "missadministrator",
					"minimal_tls_version": "1.2",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_sql_server_db_blob_auditing_policies",
					ForeignKeyName: "server_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"state":                           "Enabled",
							"retention_days":                  float64(6),
							"audit_actions_and_groups":        []interface{}{"SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP", "FAILED_DATABASE_AUTHENTICATION_GROUP", "BATCH_COMPLETED_GROUP"},
							"is_storage_secondary_key_in_use": false,
							"is_azure_monitor_target_enabled": true,
							"name":                            "Default",
							"type":                            "Microsoft.Sql/servers/auditingSettings",
						},
					}},
				},
				{
					Name:           "azure_sql_server_devops_audit_settings",
					ForeignKeyName: "server_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"is_azure_monitor_target_enabled": false,
							"state":                           "Disabled",
							"name":                            "Default",
							"type":                            "Microsoft.Sql/servers/devOpsAuditingSettings",
						},
					}},
				},
				{
					Name:           "azure_sql_server_firewall_rules",
					ForeignKeyName: "server_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"start_ip_address": "10.0.99.0",
							"end_ip_address":   "10.0.101.255",
							"name":             "office",
						},
					}},
				},
				{
					Name:           "azure_sql_server_encryption_protectors",
					ForeignKeyName: "server_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"kind":            "servicemanaged",
							"server_key_name": "ServiceManaged",
							"server_key_type": "ServiceManaged",
							"name":            "current",
							"type":            "Microsoft.Sql/servers/encryptionProtector",
						},
					}},
				},
				{
					Name:           "azure_sql_server_vulnerability_assessments",
					ForeignKeyName: "server_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"recurring_scans_is_enabled":                false,
							"recurring_scans_email_subscription_admins": true,
							"recurring_scans_emails":                    nil,
							"name":                                      "Default",
						},
					}},
				},
				{
					Name:           "azure_sql_databases",
					ForeignKeyName: "server_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":           fmt.Sprintf("mssql-db-1-%s-%s", res.Prefix, res.Suffix),
							"collation":      "SQL_Latin1_General_CP1_CI_AS",
							"license_type":   "LicenseIncluded",
							"max_size_bytes": float64(4294967296),
						},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "azure_sql_database_db_blob_auditing_policies",
							ForeignKeyName: "database_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"name": "Default",
									"type": "Microsoft.Sql/servers/databases/auditingSettings",
								},
							}},
						},
						{
							Name:           "azure_sql_database_db_threat_detection_policies",
							ForeignKeyName: "database_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"name": "default",
									"type": "Microsoft.Sql/servers/databases/securityAlertPolicies",
								},
							}},
						},
						{
							Name:           "azure_sql_database_db_vulnerability_assessments",
							ForeignKeyName: "database_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"name": "Default",
									"type": "Microsoft.Sql/servers/databases/vulnerabilityAssessments",
								},
							}},
						},
					},
				},
			},
		}
	})
}
