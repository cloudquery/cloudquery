package tableoptions

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestCloudwatchListMetricsInput_JSONSchemaExtend(t *testing.T) {
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
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput, "NextToken") + `}]}`
			}(),
		},
		{
			name: "NextToken present",
			err:  true,
			spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"list_metrics":` +
					jsonWithRemovedKeys(t, &listInput) + `}]}`
			}(),
		},
	})
}

func TestCloudwatchGetMetricStatisticsInput_JSONSchemaExtend(t *testing.T) {
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_alpha_cloudwatch_metrics":[]}`,
		},
		{
			name: "proper",
			spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"get_metric_statistics":[` +
					jsonWithRemovedKeys(t, &getInput, "Dimensions", "MetricName", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "Dimensions present",
			err:  true,
			spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"get_metric_statistics":[` +
					jsonWithRemovedKeys(t, &getInput, "MetricName", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "MetricName present",
			err:  true,
			spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"get_metric_statistics":[` +
					jsonWithRemovedKeys(t, &getInput, "Dimensions", "Namespace") + `]}]}`
			}(),
		},
		{
			name: "Namespace present",
			err:  true,
			spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `{"aws_alpha_cloudwatch_metrics":[{"get_metric_statistics":[` +
					jsonWithRemovedKeys(t, &getInput, "Dimensions", "MetricName") + `]}]}`
			}(),
		},
	})
}
