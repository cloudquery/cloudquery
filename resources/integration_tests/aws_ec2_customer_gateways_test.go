package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2CustomerGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2CustomerGateways(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2CustomerGateways().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"type":       "ipsec.1",
						"ip_address": "172.83.124.10",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ec2-cgw-%s-%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
