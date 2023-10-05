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
			Name: "proper",
			Spec: func() string {
				var listInput CloudwatchListMetricsInput
				require.NoError(t, faker.FakeObject(&listInput))
				return `[{"list_metrics":` + jsonschema.WithRemovedKeys(t, &listInput, "NextToken") + `}]`
			}(),
		},
		{
			Name: "NextToken present",
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
			Name: "proper",
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "Dimensions", "MetricName", "Namespace") + `]}]`
			}(),
		},
		{
			Name: "Dimensions present",
			Err:  true,
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "MetricName", "Namespace") + `]}]`
			}(),
		},
		{
			Name: "MetricName present",
			Err:  true,
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "Dimensions", "Namespace") + `]}]`
			}(),
		},
		{
			Name: "Namespace present",
			Err:  true,
			Spec: func() string {
				var getInput CloudwatchGetMetricStatisticsInput
				require.NoError(t, faker.FakeObject(&getInput))
				return `[{"get_metric_statistics":[` + jsonschema.WithRemovedKeys(t, &getInput, "Dimensions", "MetricName") + `]}]`
			}(),
		},
	})
}
