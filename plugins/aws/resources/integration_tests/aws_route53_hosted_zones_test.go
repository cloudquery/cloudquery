package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRoute53HostedZones(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Route53HostedZones(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_route53_hosted_zones",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("dev.%s%s.com.", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("dev.%s%s.com.", res.Prefix, res.Suffix),
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_route53_hosted_zone_resource_record_sets",
					ForeignKeyName: "hosted_zone_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name": fmt.Sprintf("dev-1.%s%s.com.dev.%s%s.com.", res.Prefix, res.Suffix, res.Prefix, res.Suffix),
						},
					}},
				},
			},
		}
	})
}
