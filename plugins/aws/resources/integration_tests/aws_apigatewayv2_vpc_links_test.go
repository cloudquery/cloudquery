package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationApigatewayv2VpcLinks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Apigatewayv2VpcLinks(), []string{"aws_apigatewayv2_vpc_links.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigatewayv2_vpc_links",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("apigw-link-%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"vpc_link_version": "V2",
					"tags": map[string]interface{}{
						"Type":   "integration_test",
						"TestId": res.Suffix,
					},
				},
			}},
		}
	})
}
