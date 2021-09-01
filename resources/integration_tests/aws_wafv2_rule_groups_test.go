package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFv2RuleGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Wafv2RuleGroups(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_wafv2_rule_groups",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("wafv2-rg-%s%s", res.Prefix, res.Suffix),
					"description": "wafv2_rule_group_1 description",
					"capacity":    float64(2),
					"policy":      nil,
					"visibility_config_cloud_watch_metrics_enabled": false,
					"visibility_config_metric_name":                 "friendly-metric-name",
					"visibility_config_sampled_requests_enabled":    false,
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
					"available_labels": []interface{}{},
					"consumed_labels":  []interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"Name": "rule-1",
							"Action": map[string]interface{}{
								"Allow": map[string]interface{}{"CustomRequestHandling": nil},
								"Block": nil,
								"Count": nil,
							},
							"Priority": float64(1),
							"Statement": map[string]interface{}{
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
							"RuleLabels":     nil,
							"OverrideAction": nil,
							"VisibilityConfig": map[string]interface{}{
								"MetricName":               "friendly-rule-metric-name",
								"SampledRequestsEnabled":   false,
								"CloudWatchMetricsEnabled": false,
							},
						},
					},
				},
			}},
		}
	})
}
