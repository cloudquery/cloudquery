package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElbv2LoadBalancers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv2LoadBalancers(), []string{"aws_elbv2_load_balancers.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elbv2_load_balancers",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("elbv2-%s", res.Suffix),
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("elbv2-%s", res.Suffix)})
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_elbv2_load_balancer_availability_zones",
					ForeignKeyName: "load_balancer_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 2,
							Data: map[string]interface{}{
								"load_balance_name": fmt.Sprintf("elbv2-%s", res.Suffix),
							},
						},
					},
				},
				{
					Name:           "aws_elbv2_load_balancer_attributes",
					ForeignKeyName: "load_balancer_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"access_logs_s3_enabled": false,
						},
					}},
				},
			},
		}
	})
}
