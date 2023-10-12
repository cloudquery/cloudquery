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

	api := Inspector2Findings{
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
	schema, err := jsonschema.Generate(Inspector2Findings{})
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
			Name: "empty list_findings",
			Spec: `{"list_findings":[]}`,
		},
		{
			Name: "null list_findings",
			Spec: `{"list_findings":null}`,
		},
		{
			Name: "bad list_findings",
			Err:  true,
			Spec: `{"list_findings":123}`,
		},
		{
			Name: "empty list_findings entry",
			Spec: `{"list_findings":[{}]}`,
		},
		{
			Name: "list_findings entry with extra keyword",
			Err:  true,
			Spec: `{"list_findings":[{"extra":123}]}`,
		},
		{
			Name: "null list_findings entry",
			Err:  true,
			Spec: `{"list_findings":[null]}`,
		},
		{
			Name: "bad list_findings entry",
			Err:  true,
			Spec: `{"list_findings":[123]}`,
		},
		{
			Name: "proper list_findings",
			Spec: func() string {
				var input CustomInspector2ListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "list_findings.NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomInspector2ListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
	})
}
