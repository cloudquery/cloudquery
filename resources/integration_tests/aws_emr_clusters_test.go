package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEmrClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EmrClusters(), []string{"aws_emr_clusters.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_emr_clusters",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("emr-cluster-%s%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"env":    "env",
							"role":   "rolename",
							"TestId": res.Suffix,
							"Type":   "integration_test",
						},
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"name": fmt.Sprintf("emr-cluster-%s%s", res.Prefix, res.Suffix)},
					squirrel.Eq{"state": "WAITING"},
				})
			},
		}
	})
}
