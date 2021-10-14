package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationPostgreSQLServers(t *testing.T) {
	table := resources.PostgresqlServers()
	const fkName = "server_cq_id"

	awsTestIntegrationHelper(t, table, []string{"azure_postgresql_servers.tf", "networks.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"administrator_login":                   "psqladminun",
					"infrastructure_encryption":             "Disabled",
					"name":                                  fmt.Sprintf("pgsqlserver-%s-%s", res.Prefix, res.Suffix),
					"public_network_access":                 "Enabled",
					"sku_name":                              "GP_Gen5_4",
					"storage_profile_backup_retention_days": float64(7),
					"storage_profile_geo_redundant_backup":  "Enabled",
					"storage_profile_storage_autogrow":      "Enabled",
					"storage_profile_storage_mb":            float64(5120),
					"version":                               "11",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_postgresql_server_configurations",
					ForeignKeyName: fkName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":  "backslash_quote",
							"value": "on",
						},
					}},
				},
				{
					Name:           "azure_postgresql_server_firewall_rules",
					ForeignKeyName: fkName,
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
					Name:           "azure_postgresql_server_private_endpoint_connections",
					ForeignKeyName: fkName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"private_link_service_connection_state_status":      "Approved",
							"private_link_service_connection_state_description": "Auto-approved",
						},
					}},
				},
			},
		}
	})
}
