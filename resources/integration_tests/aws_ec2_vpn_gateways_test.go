package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2VpnGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpnGateways(), []string{"aws_ec2_vpn_gateways.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpn_gateways",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"type": "ipsec.1",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ec2_vpn_gw_%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
