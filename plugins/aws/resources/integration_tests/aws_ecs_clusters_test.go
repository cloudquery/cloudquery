package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEcsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcsClusters(), []string{"aws_ecs_clusters.tf", "aws_vpc.tf", "aws_elbv2_load_balancers.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ecs_clusters",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix),
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix)})
			},
		}
	})
}
