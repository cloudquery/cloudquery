package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFv2ManagedRuleGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Wafv2ManagedRuleGroups(), []string{}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_wafv2_managed_rule_groups",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": "AWSManagedRulesCommonRuleSet",
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": "AWSManagedRulesAdminProtectionRuleSet",
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": "AWSManagedRulesKnownBadInputsRuleSet",
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": "AWSManagedRulesSQLiRuleSet",
					},
				},
			},
		}
	})
}
