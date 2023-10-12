package tableoptions

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListTasks(t *testing.T) {
	u := CustomECSListTasksInput{}
	require.NoError(t, faker.FakeObject(&u))
	api := ECSTasks{
		ListTasksOpts: []CustomECSListTasksInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListTasks")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListTasksOpts[0].NextToken = nil
	err = api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set Cluster in ListTasks")
}

func TestCustomECSListTasksInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(ECSTasks{})
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
			Name: "list_tasks entry with extra keyword",
			Err:  true,
			Spec: `{"list_tasks":[{"extra":123}]}`,
		},
		{
			Name: "null list_tasks entry",
			Err:  true,
			Spec: `{"list_tasks":[null]}`,
		},
		{
			Name: "bad list_tasks entry",
			Err:  true,
			Spec: `{"list_tasks":[123]}`,
		},
		{
			Name: "proper list_tasks entry",
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = aws.Int32(10) // 1-100
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken", "Cluster") + `]}`
			}(),
		},
		{
			Name: "list_tasks.NextToken present",
			Err:  true,
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = aws.Int32(10) // 1-100
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "Cluster") + `]}`
			}(),
		},
		{
			Name: "list_tasks.Cluster present",
			Err:  true,
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = aws.Int32(10) // 1-100
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "missing list_tasks.MaxResults",
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "MaxResults", "NextToken", "Cluster") + `]}`
			}(),
		},
		{
			Name: "null list_tasks.MaxResults",
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = nil
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken", "Cluster") + `]}`
			}(),
		},
		{
			Name: "zero list_tasks.MaxResults",
			Err:  true,
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = aws.Int32(0)
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken", "Cluster") + `]}`
			}(),
		},
		{
			Name: "list_tasks.MaxResults > 100",
			Err:  true,
			Spec: func() string {
				var input CustomECSListTasksInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = aws.Int32(1000)
				return `{"list_tasks":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken", "Cluster") + `]}`
			}(),
		},
	})
}
