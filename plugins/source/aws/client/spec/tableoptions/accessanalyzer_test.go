package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAAListFindings(t *testing.T) {
	u := CustomAccessAnalyzerListFindingsInput{}
	require.NoError(t, faker.FakeObject(&u))

	api := AccessanalyzerFindings{
		ListFindingOpts: []CustomAccessAnalyzerListFindingsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")
	api.ListFindingOpts[0].NextToken = nil

	err = api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set AnalyzerARN in ListFindings")
	api.ListFindingOpts[0].AnalyzerArn = nil

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	err = api.Validate()
	assert.Nil(t, err)
}

func TestCustomAccessAnalyzerListFindingsInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(AccessanalyzerFindings{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
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
			Name: "null list_findings",
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
				var input CustomAccessAnalyzerListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken", "AnalyzerArn") + `]}`
			}(),
		},
		{
			Name: "list_findings.AnalyzerArn is present",
			Err:  true,
			Spec: func() string {
				var input CustomAccessAnalyzerListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input, "AnalyzerArn") + `]}`
			}(),
		},
		{
			Name: "list_findings.NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomAccessAnalyzerListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"list_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
	})
}
