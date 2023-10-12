package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestCustomGetCostAndUsageInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(CostExplorerAPIs{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra keyword",
			Err:  true,
			Spec: `{"extra":123}`,
		},
		{
			Name: "empty get_cost_and_usage",
			Spec: `{"get_cost_and_usage":[]}`,
		},
		{
			Name: "null get_cost_and_usage",
			Spec: `{"get_cost_and_usage":null}`,
		},
		{
			Name: "bad get_cost_and_usage",
			Err:  true,
			Spec: `{"get_cost_and_usage":123}`,
		},
		{
			Name: "empty get_cost_and_usage entry",
			Spec: `{"get_cost_and_usage":[{}]}`,
		},
		{
			Name: "get_cost_and_usage entry with extra keyword",
			Err:  true,
			Spec: `{"get_cost_and_usage":[{"extra":123}]}`,
		},
		{
			Name: "null get_cost_and_usage entry",
			Err:  true,
			Spec: `{"get_cost_and_usage":[null]}`,
		},
		{
			Name: "bad get_cost_and_usage entry",
			Err:  true,
			Spec: `{"get_cost_and_usage":[123]}`,
		},
		{
			Name: "proper get_cost_and_usage",
			Spec: func() string {
				var input CustomGetCostAndUsageInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"get_cost_and_usage":[` + jsonschema.WithRemovedKeys(t, &input, "NextPageToken") + `]}`
			}(),
		},
		{
			Name: "get_cost_and_usage.NextPageToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomGetCostAndUsageInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"get_cost_and_usage":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
	})
}
