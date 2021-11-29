package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2Eips(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Eips(), []string{"aws_ec2_instances.tf", "aws_vpc.tf", "aws_ec2_eips.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2Eips().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"tags": map[string]interface{}{
							"Name":   fmt.Sprintf("elastic-ip-%s%s", res.Prefix, res.Suffix),
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
