package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRoute53HealthChecks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Route53HealthChecks(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_route53_health_checks",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("fully_qualified_domain_name = ?", fmt.Sprintf("%s%s.com", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"resource_path":    "/",
					"request_interval": 10.0,
				},
			}},
		}
	})
}
