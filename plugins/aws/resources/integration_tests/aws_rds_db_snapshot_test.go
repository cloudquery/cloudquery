package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRdsDBSnapshots(t *testing.T) {
	table := resources.RdsDbSnapshots()
	awsTestIntegrationHelper(t, table, []string{"aws_rds_instances.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"engine":                 "mysql",
						"master_username":        "foo",
						"snapshot_type":          "manual",
						"db_snapshot_identifier": fmt.Sprintf("%sdbsnap%s", res.Prefix, res.Suffix),
						"attributes": []interface{}{
							map[string]interface{}{
								"AttributeName":   "restore",
								"AttributeValues": []interface{}{},
							},
						},
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
