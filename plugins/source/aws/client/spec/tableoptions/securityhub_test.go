package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFindings(t *testing.T) {
	u := CustomGetFindingsOpts{}
	require.NoError(t, faker.FakeObject(&u))
	api := SecurityHubAPIs{
		GetFindingsOpts: []CustomGetFindingsOpts{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in GetFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.GetFindingsOpts[0].NextToken = nil
	err = api.Validate()
	assert.EqualError(t, err, "invalid range: MaxResults must be within range [1-100]")
}

func TestCustomGetFindingsOpts_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(SecurityHubAPIs{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 10 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 10 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
		{
			Name: "MaxResults > 100",
			Err:  true,
			Spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 1000 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "MaxResults < 1",
			Err:  true,
			Spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 0 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
	})
}
