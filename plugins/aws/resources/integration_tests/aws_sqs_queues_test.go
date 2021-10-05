package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSQSQueues(t *testing.T) {
	table := resources.SQSQueues()
	awsTestIntegrationHelper(t, table, []string{"aws_sns.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"fifo_queue":                  true,
					"content_based_deduplication": true,
					"deduplication_scope":         "queue",
					"fifo_throughput_limit":       "perQueue",
					"tags": map[string]interface{}{
						"Type":   "integration_test",
						"TestId": res.Suffix,
					},
				},
			}},
		}
	})
}
