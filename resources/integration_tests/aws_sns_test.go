package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSnsSubscriptions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SnsSubscriptions(), []string{"aws_sns.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_sns_subscriptions",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"protocol": "sqs",
					},
				},
			},
			Filter: func(sb squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sb.Where(squirrel.ILike{"arn": fmt.Sprintf("%%%s%%", res.Suffix)})
			},
		}
	})
}

func TestIntegrationSnsTopics(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SnsTopics(), []string{"aws_sns.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_sns_topics",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"content_based_deduplication": true,
						"display_name":                fmt.Sprintf("%s-%s", res.Prefix, res.Suffix),
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"content_based_deduplication": false,
						"display_name":                fmt.Sprintf("sns-test2-%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
			Filter: func(sb squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sb.Where(squirrel.Or{
					squirrel.Eq{"display_name": fmt.Sprintf("%s-%s", res.Prefix, res.Suffix)},
					squirrel.Eq{"display_name": fmt.Sprintf("sns-test2-%s-%s", res.Prefix, res.Suffix)},
				})
			},
		}
	})
}
