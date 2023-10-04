package tableoptions

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestCustomGetCostAndUsageInput_JSONSchemaExtend(t *testing.T) {
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_alpha_costexplorer_cost_custom":{}}`,
		},
		{
			name: "proper",
			spec: func() string {
				var input CustomGetCostAndUsageInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_alpha_costexplorer_cost_custom":{"get_cost_and_usage":[` +
					jsonWithRemovedKeys(t, &input, "NextPageToken") + `]}}`
			}(),
		},
		{
			name: "NextPageToken is present",
			err:  true,
			spec: func() string {
				var input CustomGetCostAndUsageInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_alpha_costexplorer_cost_custom":{"get_cost_and_usage":[` +
					jsonWithRemovedKeys(t, &input) + `]}}`
			}(),
		},
	})
}
