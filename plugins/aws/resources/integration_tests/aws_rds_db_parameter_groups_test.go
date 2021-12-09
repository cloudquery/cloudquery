package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRdsDbParameterGroups(t *testing.T) {
	table := resources.RdsDbParameterGroups()
	awsTestIntegrationHelper(t, table, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":        fmt.Sprintf("rds-db-pg-%s-%s", res.Prefix, res.Suffix),
						"family":      "mysql8.0",
						"description": "Test RDS DB parameter group",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{{
				Name:           "aws_rds_db_parameters",
				ForeignKeyName: "db_parameter_group_cq_id",
				ExpectedValues: []providertest.ExpectedValue{
					{
						Count: 1,
						Data: map[string]interface{}{
							"parameter_name":  "character_set_client",
							"parameter_value": "utf8",
						},
					},
					{
						Count: 1,
						Data: map[string]interface{}{
							"parameter_name":  "character_set_server",
							"parameter_value": "utf8",
						},
					},
				},
			}},
		}
	})
}
