package tableoptions

import (
	"testing"

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
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_securityhub_findings":{}}`,
		},
		{
			name: "proper",
			spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 10 // range 1-100
				return `{"aws_securityhub_findings":{"get_findings":[` +
					jsonWithRemovedKeys(t, &input, "NextToken") + `]}}`
			}(),
		},
		{
			name: "NextToken is present",
			err:  true,
			spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 10 // range 1-100
				return `{"aws_securityhub_findings":{"get_findings":[` +
					jsonWithRemovedKeys(t, &input) + `]}}`
			}(),
		},
		{
			name: "MaxResults > 100",
			err:  true,
			spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 1000 // range 1-100
				return `{"aws_securityhub_findings":{"get_findings":[` +
					jsonWithRemovedKeys(t, &input, "NextToken") + `]}}`
			}(),
		},
		{
			name: "MaxResults < 1",
			err:  true,
			spec: func() string {
				var input CustomGetFindingsOpts
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 0 // range 1-100
				return `{"aws_securityhub_findings":{"get_findings":[` +
					jsonWithRemovedKeys(t, &input, "NextToken") + `]}}`
			}(),
		},
	})
}
