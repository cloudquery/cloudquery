package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationMonitoringAlertPolicies(t *testing.T) {
	testIntegrationHelper(t, resources.MonitoringAlertPolicies(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.MonitoringAlertPolicies().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"display_name": fmt.Sprintf("alert-policies-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"combiner":     "OR",
						"enabled":      true,
						"display_name": fmt.Sprintf("alert-policies-%s-%s", res.Prefix, res.Suffix),
						"labels": map[string]interface{}{
							"foo": "bar",
						},
						"validity_code": float64(0),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_monitoring_alert_policy_conditions",
					ForeignKeyName: "alert_policy_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"absent_trigger_count":                      float64(0),
								"absent_trigger_percent":                    float64(0),
								"monitoring_query_language_trigger_count":   float64(0),
								"monitoring_query_language_trigger_percent": float64(0),
								"threshold_value":                           0.6,
								"threshold_trigger_count":                   float64(0),
								"threshold_trigger_percent":                 float64(0),
								"display_name":                              "StatefulSet has enough ready replicas",
								"threshold_filter":                          fmt.Sprintf("metric.type=\"logging.googleapis.com/user/alerts-metric-%s-%s\" resource.type=\"gke_container\"", res.Prefix, res.Suffix),
								"threshold_denominator_filter":              fmt.Sprintf("metric.type=\"logging.googleapis.com/user/alerts-metric1-%s-%s\" resource.type=\"gke_container\"", res.Prefix, res.Suffix),
								"threshold_duration":                        "300s",
								"threshold_comparison":                      "COMPARISON_LT",
								"absent_duration":                           "",
							},
						},
					},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "gcp_monitoring_alert_policy_condition_threshold_aggregations",
							ForeignKeyName: "alert_policy_condition_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"alignment_period":     "60s",
										"per_series_aligner":   "",
										"cross_series_reducer": "",
									},
								},
							},
						},
						{
							Name:           "gcp_monitoring_alert_policy_condition_denominator_aggs",
							ForeignKeyName: "alert_policy_condition_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"alignment_period":     "60s",
										"per_series_aligner":   "",
										"cross_series_reducer": "",
									},
								},
							},
						},
					},
				},
			},
		}
	})
}

func TestIntegrationMonitoringAlertPolicies_absent(t *testing.T) {
	testIntegrationHelper(t, resources.MonitoringAlertPolicies(), []string{
		"gcp_monitoring_alert_policies_absent.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.MonitoringAlertPolicies().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"display_name": fmt.Sprintf("alert-absent-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"combiner":     "OR",
						"enabled":      true,
						"display_name": fmt.Sprintf("alert-absent-%s-%s", res.Prefix, res.Suffix),
						"labels": map[string]interface{}{
							"foo": "bar",
						},
						"validity_code": float64(0),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_monitoring_alert_policy_conditions",
					ForeignKeyName: "alert_policy_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"absent_trigger_count":                      float64(0),
								"absent_trigger_percent":                    float64(0),
								"monitoring_query_language_trigger_count":   float64(0),
								"monitoring_query_language_trigger_percent": float64(0),
								"threshold_value":                           float64(0),
								"threshold_trigger_count":                   float64(0),
								"threshold_trigger_percent":                 float64(0),
								"display_name":                              "test condition1",
								"threshold_filter":                          "",
								"threshold_denominator_filter":              "",
								"threshold_duration":                        "",
								"threshold_comparison":                      "",
								"absent_duration":                           "120s",
								"absent_filter":                             fmt.Sprintf("metric.type=\"logging.googleapis.com/user/alerts-absent-metric-%s-%s\" AND resource.type=\"metric\"", res.Prefix, res.Suffix),
							},
						},
					},
					Relations: []*providertest.ResourceIntegrationVerification{

						{
							Name:           "gcp_monitoring_alert_policy_condition_absent_aggregations",
							ForeignKeyName: "alert_policy_condition_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"alignment_period":   "120s",
										"per_series_aligner": "ALIGN_RATE",
									},
								},
							},
						},
					},
				},
			},
		}
	})
}
