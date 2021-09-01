package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationApigatewayVpcLinks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayVpcLinks(), []string{"aws_apigateway.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigateway_vpc_links",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("apigw-vpc-link-%s-%s", res.Prefix, res.Suffix),
					"description": "example description",
				},
			}},
		}
	})
}
