package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCloudwatchlogsFilters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudwatchlogsFilters(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_cloudwatchlogs_filters",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":           fmt.Sprintf("aws_cloudwatch_log_metric_filter_%s%s", res.Prefix, res.Suffix),
						"log_group_name": fmt.Sprintf("MyApp%s%s/access.log", res.Prefix, res.Suffix),
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("aws_cloudwatch_log_metric_filter_%s%s", res.Prefix, res.Suffix)})
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_cloudwatchlogs_filter_metric_transformations",
					ForeignKeyName: "filter_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"metric_name":      "aws_cloudwatch_log_metric_filter_name",
								"metric_namespace": "YourNamespace",
								"metric_value":     "1",
							},
						},
					},
				},
			},
		}
	})
}
