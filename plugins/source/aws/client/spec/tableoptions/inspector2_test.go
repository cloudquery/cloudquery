package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInspector2ListFindings(t *testing.T) {
	u := CustomInspector2ListFindingsInput{}
	require.NoError(t, faker.FakeObject(&u))

	api := Inspector2APIs{
		ListFindingsOpts: []CustomInspector2ListFindingsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListFindingsOpts[0].NextToken = nil
	err = api.Validate()

	assert.Nil(t, err)
}

func TestCustomInspector2ListFindingsInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(Inspector2APIs{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input CustomInspector2ListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomInspector2ListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
	})
}
