package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2EbsVolumes(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2EbsVolumes(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2EbsVolumes().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"volume_type":          "gp2",
						"multi_attach_enabled": false,
						"encrypted":            false,
						"size":                 float64(5),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ec2-ebs-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
