package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFWebACLs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafWebAcls(), []string{"aws_waf_rules.tf", "aws_waf_web_acls.tf", "aws_kinesis_firehose.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_waf_web_acls",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                fmt.Sprintf("waf-web-acl-%s%s", res.Prefix, res.Suffix),
					"metric_name":         "wafwebacl1",
					"default_action_type": "ALLOW",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{{
				Name:           "aws_waf_web_acl_rules",
				ForeignKeyName: "web_acl_cq_id",
				ExpectedValues: []providertest.ExpectedValue{{
					Count: 1,
					Data: map[string]interface{}{
						"priority":    float64(1),
						"action_type": "BLOCK",
						"type":        "REGULAR",
					},
				}},
			}},
		}
	})
}
