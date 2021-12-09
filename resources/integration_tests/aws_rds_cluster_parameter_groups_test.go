package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRdsClusterParameterGroups(t *testing.T) {
	table := resources.RdsClusterParameterGroups()
	awsTestIntegrationHelper(t, table, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":        fmt.Sprintf("rds-cluster-pg-%s-%s", res.Prefix, res.Suffix),
						"family":      "aurora-mysql8.0",
						"description": "Test RDS cluster parameter group",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{{
				Name:           "aws_rds_cluster_parameters",
				ForeignKeyName: "cluster_parameter_group_cq_id",
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
