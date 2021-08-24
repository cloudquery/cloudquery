package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDirectCloudwatchAlarms(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudwatchAlarms(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_cloudwatch_alarms",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("cl-alarm%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"actions_enabled": true,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_cloudwatch_alarm_metrics",
					ForeignKeyName: "alarm_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"id": "m1",
						},
					}},
				},
				{
					Name:           "aws_cloudwatch_alarm_metrics",
					ForeignKeyName: "alarm_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"id": "m2",
						},
					}},
				},
				{
					Name:           "aws_cloudwatch_alarm_metrics",
					ForeignKeyName: "alarm_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"id":    "e1",
							"label": "Error Rate",
						},
					}},
				},
			},
		}
	})
}
