package tableoptions

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestCloudwatchMetricsJSONSchema(t *testing.T) {
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_alpha_cloudwatch_metrics":[]}`,
		},
		{
			name: "proper",
			spec: func() string {
				var metrics CloudwatchMetric
				require.NoError(t, faker.FakeObject(&metrics))

				// remove prohibited fields
				metrics.ListMetricsOpts.NextToken = nil
				require.Len(t, metrics.GetMetricStatisticsOpts, 1)
				metrics.GetMetricStatisticsOpts[0].Dimensions = nil
				metrics.GetMetricStatisticsOpts[0].MetricName = nil
				metrics.GetMetricStatisticsOpts[0].Namespace = nil

				data, err := json.MarshalIndent(TableOptions{CloudwatchMetrics: CloudwatchMetrics{metrics}}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Namespace\": null,\n", ``, 1)
				result = strings.Replace(result, "\"MetricName\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Dimensions\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "list_metrics.NextToken present",
			err:  true,
			spec: func() string {
				var metrics CloudwatchMetric
				require.NoError(t, faker.FakeObject(&metrics))

				// remove prohibited fields
				require.Len(t, metrics.GetMetricStatisticsOpts, 1)
				metrics.GetMetricStatisticsOpts[0].Dimensions = nil
				metrics.GetMetricStatisticsOpts[0].MetricName = nil
				metrics.GetMetricStatisticsOpts[0].Namespace = nil

				data, err := json.MarshalIndent(TableOptions{CloudwatchMetrics: CloudwatchMetrics{metrics}}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"Namespace\": null,\n", ``, 1)
				result = strings.Replace(result, "\"MetricName\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Dimensions\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "get_metric_statistics.Dimensions present",
			err:  true,
			spec: func() string {
				var metrics CloudwatchMetric
				require.NoError(t, faker.FakeObject(&metrics))

				// remove prohibited fields
				metrics.ListMetricsOpts.NextToken = nil
				require.Len(t, metrics.GetMetricStatisticsOpts, 1)
				metrics.GetMetricStatisticsOpts[0].MetricName = nil
				metrics.GetMetricStatisticsOpts[0].Namespace = nil

				data, err := json.MarshalIndent(TableOptions{CloudwatchMetrics: CloudwatchMetrics{metrics}}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Namespace\": null,\n", ``, 1)
				result = strings.Replace(result, "\"MetricName\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "get_metric_statistics.MetricName present",
			err:  true,
			spec: func() string {
				var metrics CloudwatchMetric
				require.NoError(t, faker.FakeObject(&metrics))

				// remove prohibited fields
				metrics.ListMetricsOpts.NextToken = nil
				require.Len(t, metrics.GetMetricStatisticsOpts, 1)
				metrics.GetMetricStatisticsOpts[0].Dimensions = nil
				metrics.GetMetricStatisticsOpts[0].Namespace = nil

				data, err := json.MarshalIndent(TableOptions{CloudwatchMetrics: CloudwatchMetrics{metrics}}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Namespace\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Dimensions\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "get_metric_statistics.Namespace present",
			err:  true,
			spec: func() string {
				var metrics CloudwatchMetric
				require.NoError(t, faker.FakeObject(&metrics))

				// remove prohibited fields
				metrics.ListMetricsOpts.NextToken = nil
				require.Len(t, metrics.GetMetricStatisticsOpts, 1)
				metrics.GetMetricStatisticsOpts[0].Dimensions = nil
				metrics.GetMetricStatisticsOpts[0].MetricName = nil

				data, err := json.MarshalIndent(TableOptions{CloudwatchMetrics: CloudwatchMetrics{metrics}}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				result = strings.Replace(result, "\"MetricName\": null,\n", ``, 1)
				result = strings.Replace(result, "\"Dimensions\": null,\n", ``, 1)
				return result
			}(),
		},
	})
}
