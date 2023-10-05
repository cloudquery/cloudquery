package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupEvents(t *testing.T) {
	u := CustomLookupEventsOpts{}
	require.NoError(t, faker.FakeObject(&u))

	api := CloudtrailAPIs{
		LookupEventsOpts: []CustomLookupEventsOpts{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in LookupEvents")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.LookupEventsOpts[0].NextToken = nil
	err = api.Validate()
	assert.Nil(t, err)
}

func TestCustomLookupEventsOpts_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{"aws_cloudtrail_events":{}}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input CustomLookupEventsOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_cloudtrail_events":{"lookup_events":[` +
					jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}}`
			}(),
		},
		{
			Name: "NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomLookupEventsOpts
				require.NoError(t, faker.FakeObject(&input))
				return `{"aws_cloudtrail_events":{"lookup_events":[` +
					jsonschema.WithRemovedKeys(t, &input) + `]}}`
			}(),
		},
	})
}
