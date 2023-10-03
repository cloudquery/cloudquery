package tableoptions

import (
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
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput, "NextToken") + `,` +
					`"get_metric_statistics":[` + jsonWithRemovedKeys(t, &getInput, "Dimensions", "MetricName", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "list_metrics.NextToken present",
			err:  true,
			spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput) + `,` +
					`"get_metric_statistics":[` + jsonWithRemovedKeys(t, &getInput, "Dimensions", "MetricName", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "get_metric_statistics.Dimensions present",
			err:  true,
			spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput, "NextToken") + `,` +
					`"get_metric_statistics":[` + jsonWithRemovedKeys(t, &getInput, "MetricName", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "get_metric_statistics.MetricName present",
			err:  true,
			spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput, "NextToken") + `,` +
					`"get_metric_statistics":[` + jsonWithRemovedKeys(t, &getInput, "Dimensions", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "get_metric_statistics.Namespace present",
			err:  true,
			spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput, "NextToken") + `,` +
					`"get_metric_statistics":[` + jsonWithRemovedKeys(t, &getInput, "Dimensions", "MetricName") + `]}]}`
			}(),
		},
	})
}
