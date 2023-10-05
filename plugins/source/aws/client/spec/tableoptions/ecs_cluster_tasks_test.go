package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListTasks(t *testing.T) {
	u := CustomListTasksOpts{}
	require.NoError(t, faker.FakeObject(&u))
	api := ECSTaskAPIs{
		ListTasksOpts: []CustomListTasksOpts{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListTasks")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListTasksOpts[0].NextToken = nil
	err = api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set Cluster in ListTasks")
}

func TestCustomListTasksOpts_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{"aws_ecs_cluster_tasks":{}}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input CustomListTasksOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_ecs_cluster_tasks":{"list_tasks":[` +
					jsonschema.WithRemovedKeys(t, &input, "NextToken", "Cluster") + `]}}`
			}(),
		},
		{
			Name: "NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomListTasksOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_ecs_cluster_tasks":{"list_tasks":[` +
					jsonschema.WithRemovedKeys(t, &input, "Cluster") + `]}}`
			}(),
		},
		{
			Name: "Cluster is present",
			Err:  true,
			Spec: func() string {
				var input CustomListTasksOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_ecs_cluster_tasks":{"list_tasks":[` +
					jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}}`
			}(),
		},
	})
}
