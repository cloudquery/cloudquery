package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSQLInstances(t *testing.T) {
	const fkName = "instance_cq_id"
	table := resources.SQLInstances()
	testIntegrationHelper(t, table, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":             fmt.Sprintf("sql-database-inst-%s%s-v2", res.Prefix, res.Suffix),
					"database_version": "POSTGRES_11",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_sql_instance_ip_addresses",
					ForeignKeyName: fkName,
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"type": "PRIMARY",
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"type": "OUTGOING",
							},
						},
					},
				},
				{
					Name:           "gcp_sql_instance_settings_ip_config_authorized_networks",
					ForeignKeyName: fkName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":  "testnet",
							"value": "8.8.8.8",
						},
					}},
				},
			},
		}
	})
}
