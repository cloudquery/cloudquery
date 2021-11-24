package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2EbsSnapshots(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2EbsSnapshots(), []string{"aws_ec2_ebs_snapshots.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2EbsSnapshots().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						// "volume_size": 40,
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ec2-ebs-snapshot-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
