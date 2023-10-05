package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestCloudwatchListMetricsInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(CloudwatchMetrics{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `[]`,
		},
		{
			Name: "empty list_metrics",
			Spec: `[{"list_metrics":{}}]`,
		},
		{
			Name: "null list_metrics",
			Err:  true,
			Spec: `[{"list_metrics":null}]`,
		},
		{
			Name: "bad list_metrics",
			Err:  true,
			Spec: `[{"list_metrics":123}]`,
		},
		{
			Name: "proper list_metrics",
			Spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				return `[{"list_metrics":` + jsonschema.WithRemovedKeys(t, &listInput, "NextToken") + `}]`
			}(),
		},
		{
			Name: "list_metrics.NextToken present",
			Err:  true,
			Spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				return `[{"list_metrics":` + jsonschema.WithRemovedKeys(t, &listInput) + `}]`
			}(),
		},
	})
}

func TestCloudwatchGetMetricStatisticsInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(CloudwatchMetrics{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `[]`,
		},
		{
			Name: "empty get_metric_statistics",
			Spec: `[{"get_metric_statistics":[]}]`,
		},
		{
			Name: "null get_metric_statistics",
			Spec: `[{"get_metric_statistics":null}]`,
		},
		{
			Name: "bad get_metric_statistics",
			Err:  true,
			Spec: `[{"get_metric_statistics":123}]`,
		},
		{
			Name: "empty get_metric_statistics entry",
			Spec: `[{"get_metric_statistics":[{}]}]`,
		},
		{
			Name: "null get_metric_statistics entry",
			Err:  true,
			Spec: `[{"get_metric_statistics":[null]}]`,
		},
		{
			Name: "bad get_metric_statistics entry",
			Err:  true,
			Spec: `[{"get_metric_statistics":[123]}]`,
		},
		{
			Name: "proper get_metric_statistics",
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "Dimensions", "MetricName", "Namespace") + `]}]`
			}(),
		},
		{
			Name: "get_metric_statistics.Dimensions present",
			Err:  true,
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "MetricName", "Namespace") + `]}]`
			}(),
		},
		{
			Name: "get_metric_statistics.MetricName present",
			Err:  true,
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "Dimensions", "Namespace") + `]}]`
			}(),
		},
		{
			Name: "get_metric_statistics.Namespace present",
			Err:  true,
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "Dimensions", "MetricName") + `]}]`
			}(),
		},
	})
}
