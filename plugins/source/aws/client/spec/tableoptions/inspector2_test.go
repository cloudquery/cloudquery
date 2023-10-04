package tableoptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cloudquery/plugin-sdk/faker"
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
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_inspector2_findings":{}}`,
		},
		{
			name: "proper",
			spec: func() string {
				var input CustomInspector2ListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_inspector2_findings":{"list_findings":[` +
					jsonWithRemovedKeys(t, &input, "NextToken") + `]}}`
			}(),
		},
		{
			name: "NextToken is present",
			err:  true,
			spec: func() string {
				var input CustomInspector2ListFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_inspector2_findings":{"list_findings":[` +
					jsonWithRemovedKeys(t, &input) + `]}}`
			}(),
		},
	})
}
