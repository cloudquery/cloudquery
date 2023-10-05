package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestCustomGetCostAndUsageInput_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{"aws_alpha_costexplorer_cost_custom":{}}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input CustomGetCostAndUsageInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_alpha_costexplorer_cost_custom":{"get_cost_and_usage":[` +
					jsonschema.WithRemovedKeys(t, &input, "NextPageToken") + `]}}`
			}(),
		},
		{
			Name: "NextPageToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomGetCostAndUsageInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_alpha_costexplorer_cost_custom":{"get_cost_and_usage":[` +
					jsonschema.WithRemovedKeys(t, &input) + `]}}`
			}(),
		},
	})
}
