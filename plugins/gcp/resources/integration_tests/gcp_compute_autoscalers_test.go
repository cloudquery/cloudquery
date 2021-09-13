package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeAutoscalers(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeAutoscalers(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeAutoscalers().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("autoscaler%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                             fmt.Sprintf("autoscaler%s%s", res.Prefix, res.Suffix),
						"kind":                             "compute#autoscaler",
						"scale_in_control_time_window_sec": float64(0),
						"scale_in_control_max_scaled_in_replicas_percent":    float64(0),
						"scale_in_control_max_scaled_in_replicas_fixed":      float64(0),
						"scale_in_control_max_scaled_in_replicas_calculated": float64(0),
						"mode":             "ON",
						"min_num_replicas": float64(1),
						"max_num_replicas": float64(5),
						"load_balancing_utilization_utilization_target": float64(0),
						"cool_down_period_sec":                          float64(60),
						"cpu_utilization_predictive_method":             "NONE",
						"cpu_utilization_utilization_target":            0.5,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_autoscaler_custom_metric_utilizations",
					ForeignKeyName: "autoscaler_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"metric":                     "pubsub.googleapis.com/subscription/num_undelivered_messages",
								"single_instance_assignment": float64(0),
								"utilization_target":         float64(100),
								"utilization_target_type":    "DELTA_PER_MINUTE",
							},
						},
					},
				},
			},
		}
	})
}
