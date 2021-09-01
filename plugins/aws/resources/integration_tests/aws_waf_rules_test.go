package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFRules(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafRules(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_waf_rules",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("waf-rule-%s%s", res.Prefix, res.Suffix),
					"metric_name": "wafrule1",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{{
				Name:           "aws_waf_rule_predicates",
				ForeignKeyName: "rule_cq_id",
				ExpectedValues: []providertest.ExpectedValue{{
					Count: 1,
					Data: map[string]interface{}{
						"negated": false,
						"type":    "IPMatch",
					},
				}},
			}},
		}
	})
}
