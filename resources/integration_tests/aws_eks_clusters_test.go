package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEksClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EksClusters(), []string{"aws_eks_clusters.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_eks_clusters",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("eks-%s%s", res.Prefix, res.Suffix),
						"resources_vpc_config_endpoint_private_access": false,
						"resources_vpc_config_endpoint_public_access":  true,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_eks_cluster_loggings",
					ForeignKeyName: "cluster_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"enabled": false,
						},
					}},
				},
			},
		}
	})
}
