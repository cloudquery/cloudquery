package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCloudtrailTrails(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudtrailTrails(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_cloudtrail_trails",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("cloudtrail-%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                          fmt.Sprintf("cloudtrail-%s-%s", res.Prefix, res.Suffix),
					"s3_key_prefix":                 "cloudtrail",
					"include_global_service_events": true,
					"is_multi_region_trail":         true,
					"is_organization_trail":         false,
					"log_file_validation_enabled":   true,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_cloudtrail_trail_event_selectors",
					ForeignKeyName: "trail_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"include_management_events": true,
							"read_write_type":           "All",
						},
					}},
				},
			},
		}
	})
}
