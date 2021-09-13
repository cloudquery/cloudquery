package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationLoggingMetrics(t *testing.T) {
	testIntegrationHelper(t, resources.LoggingMetrics(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.LoggingMetrics().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("logging-metrics-metric-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                                           fmt.Sprintf("logging-metrics-metric-%s%s", res.Prefix, res.Suffix),
						"filter":                                         "protoPayload.methodName=\"cloudsql.instances.update\"",
						"metric_descriptor_value_type":                   "DISTRIBUTION",
						"metric_descriptor_metric_kind":                  "DELTA",
						"linear_buckets_options_width":                   float64(1),
						"linear_buckets_options_offset":                  float64(1),
						"linear_buckets_options_num_finite_buckets":      float64(3),
						"exponential_buckets_options_scale":              float64(0),
						"exponential_buckets_options_num_finite_buckets": float64(0),
						"exponential_buckets_options_growth_factor":      float64(0),
						"label_extractors": map[string]interface{}{
							"sku":  "EXTRACT(jsonPayload.id)",
							"mass": "EXTRACT(jsonPayload.request)",
						},
						"metric_descriptor_display_name": "My metric",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_logging_metric_descriptor_labels",
					ForeignKeyName: "metric_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"description": "Identifying number for item",
								"key":         "sku",
								"value_type":  "INT64",
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"description": "amount of matter",
								"key":         "mass",
							},
						},
					},
				},
			},
		}
	})
}
