package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRdsEventSubscriptions(t *testing.T) {
	table := resources.RdsEventSubscriptions()
	awsTestIntegrationHelper(t, table, []string{"aws_rds_event_subscriptions.tf", "aws_sns.tf", "aws_rds_instances.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"cust_subscription_id":  fmt.Sprintf("rds-event-sub-%s-%s", res.Prefix, res.Suffix),
						"enabled":               true,
						"event_categories_list": []interface{}{"failure"},
						"source_type":           "db-instance",
						"status":                "active",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
