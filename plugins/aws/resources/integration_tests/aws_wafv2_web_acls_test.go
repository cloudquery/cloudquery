package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFv2WebACLs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Wafv2WebAcls(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_wafv2_web_acls",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("wafv2-web-acl-%s%s", res.Prefix, res.Suffix),
					"description": "Example of a managed rule.",
					"default_action": map[string]interface{}{
						"Allow": map[string]interface{}{
							"CustomRequestHandling": nil,
						},
						"Block": nil,
					},
					"visibility_config_cloud_watch_metrics_enabled": false,
					"visibility_config_metric_name":                 "friendly-metric-name",
					"visibility_config_sampled_requests_enabled":    false,
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
						"Tag1":   "Value1",
						"Tag2":   "Value2",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_wafv2_web_acl_rules",
					ForeignKeyName: "web_acl_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":     "rule-1",
							"priority": float64(1),
							"visibility_config_cloud_watch_metrics_enabled": false,
							"visibility_config_metric_name":                 "friendly-rule-metric-name",
							"visibility_config_sampled_requests_enabled":    false,
							"override_action": map[string]interface{}{
								"None": nil,
								"Count": map[string]interface{}{
									"CustomRequestHandling": nil,
								},
							},
							"statement": map[string]interface{}{
								"OrStatement":             nil,
								"AndStatement":            nil,
								"NotStatement":            nil,
								"GeoMatchStatement":       nil,
								"XssMatchStatement":       nil,
								"ByteMatchStatement":      nil,
								"RateBasedStatement":      nil,
								"SqliMatchStatement":      nil,
								"LabelMatchStatement":     nil,
								"IPSetReferenceStatement": nil,
								"SizeConstraintStatement": nil,
								"ManagedRuleGroupStatement": map[string]interface{}{
									"Name":       "AWSManagedRulesCommonRuleSet",
									"VendorName": "AWS",
									"ExcludedRules": []interface{}{
										map[string]interface{}{"Name": "SizeRestrictions_QUERYSTRING"},
										map[string]interface{}{"Name": "NoUserAgent_HEADER"},
									},
									"ScopeDownStatement": map[string]interface{}{
										"OrStatement":  nil,
										"AndStatement": nil,
										"NotStatement": nil,
										"GeoMatchStatement": map[string]interface{}{
											"CountryCodes":      []interface{}{"US", "NL"},
											"ForwardedIPConfig": nil,
										},
										"XssMatchStatement":                 nil,
										"ByteMatchStatement":                nil,
										"RateBasedStatement":                nil,
										"SqliMatchStatement":                nil,
										"LabelMatchStatement":               nil,
										"IPSetReferenceStatement":           nil,
										"SizeConstraintStatement":           nil,
										"ManagedRuleGroupStatement":         nil,
										"RuleGroupReferenceStatement":       nil,
										"RegexPatternSetReferenceStatement": nil,
									},
								},
								"RuleGroupReferenceStatement":       nil,
								"RegexPatternSetReferenceStatement": nil,
							},
						},
					}},
				},
			},
		}
	})
}
