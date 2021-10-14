package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationMySQLServers(t *testing.T) {
	table := resources.MySQLServers()
	const fkName = "server_cq_id"

	awsTestIntegrationHelper(t, table, []string{"azure_mysql_servers.tf", "networks.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"administrator_login":                   "mysqladminun",
					"infrastructure_encryption":             "Disabled",
					"name":                                  fmt.Sprintf("mysql-server-%s-%s", res.Prefix, res.Suffix),
					"public_network_access":                 "Enabled",
					"sku_name":                              "GP_Gen5_2",
					"storage_profile_backup_retention_days": float64(7),
					"storage_profile_geo_redundant_backup":  "Disabled",
					"storage_profile_storage_autogrow":      "Enabled",
					"storage_profile_storage_mb":            float64(5120),
					"version":                               "5.7",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_mysql_server_configurations",
					ForeignKeyName: fkName,
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":  "interactive_timeout",
								"value": "600",
							},
						},
					},
				},
				{
					Name:           "azure_mysql_server_private_endpoint_connections",
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
