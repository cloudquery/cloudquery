package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRdsClusterSnapshots(t *testing.T) {
	table := resources.RdsClusterSnapshots()
	awsTestIntegrationHelper(t, table, []string{"aws_rds_clusters.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"db_cluster_snapshot_identifier": fmt.Sprintf("%ssnap%s", res.Prefix, res.Suffix),
						"db_cluster_identifier":          fmt.Sprintf("rdscluster%s", res.Suffix),
						"engine":                         "aurora",
						"master_username":                "foo",
						"snapshot_type":                  "manual",
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
