package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRedshiftSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RedshiftSubnetGroups(), []string{"aws_vpc.tf", "aws_redshift_subnet_groups.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_redshift_subnet_groups",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"cluster_subnet_group_name": fmt.Sprintf("redshift-sg-%s%s", res.Prefix, res.Suffix),
					"description":               "my test description",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_redshift_subnet_group_subnets",
					ForeignKeyName: "subnet_group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 2,
					}},
				},
			},
		}
	})
}
