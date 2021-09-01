package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFRuleGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafRuleGroups(), []string{"aws_waf_rules.tf", "aws_waf_rule_groups.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_waf_rule_groups",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("waf-rg-%s%s", res.Prefix, res.Suffix),
					"metric_name": "wafrulegroup1",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
		}
	})
}
