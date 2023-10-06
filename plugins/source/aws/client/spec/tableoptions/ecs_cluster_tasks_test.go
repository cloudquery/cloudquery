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
	schema, err := jsonschema.Generate(ECSTaskAPIs{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "empty list_tasks",
			Spec: `{"list_tasks":[]}`,
		},
		{
			Name: "null list_tasks",
			Spec: `{"list_tasks":null}`,
		},
		{
			Name: "bad list_tasks",
			Err:  true,
			Spec: `{"list_tasks":123}`,
		},
		{
			Name: "empty list_tasks entry",
			Spec: `{"list_tasks":[{}]}`,
		},
		{
			Name: "null list_tasks entry",
			Err:  true,
			Spec: `{"list_tasks":[null]}`,
		},
		{
			Name: "bad list_tasks entry",
			Err:  true,
			Spec: `{"list_tasks":123}`,
		},
		{
			Name: "proper list_tasks",
			Spec: func() string {
				var input CustomListTasksOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken", "Cluster") + `]}`
			}(),
		},
		{
			Name: "list_tasks.NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomListTasksOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "Cluster") + `]}`
			}(),
		},
		{
			Name: "list_tasks.Cluster is present",
			Err:  true,
			Spec: func() string {
				var input CustomListTasksOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
	})
}
