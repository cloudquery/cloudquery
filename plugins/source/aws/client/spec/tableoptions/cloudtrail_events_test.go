package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupEvents(t *testing.T) {
	u := CustomCloudtrailLookupEventsInput{}
	require.NoError(t, faker.FakeObject(&u))

	api := CloudtrailEvents{
		LookupEventsOpts: []CustomCloudtrailLookupEventsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in LookupEvents")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.LookupEventsOpts[0].NextToken = nil
	err = api.Validate()
	assert.Nil(t, err)
}

func TestCustomCloudtrailLookupEventsInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(CloudtrailEvents{})
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
			Name: "empty lookup_events",
			Spec: `{"lookup_events":[]}`,
		},
		{
			Name: "null lookup_events",
			Spec: `{"lookup_events":null}`,
		},
		{
			Name: "bad lookup_events",
			Err:  true,
			Spec: `{"lookup_events":123}`,
		},
		{
			Name: "empty lookup_events entry",
			Spec: `{"lookup_events":[{}]}`,
		},
		{
			Name: "lookup_events entry with extra keyword",
			Err:  true,
			Spec: `{"lookup_events":[{"extra":123}]}`,
		},
		{
			Name: "null lookup_events entry",
			Err:  true,
			Spec: `{"lookup_events":[null]}`,
		},
		{
			Name: "bad lookup_events entry",
			Err:  true,
			Spec: `{"lookup_events":[123]}`,
		},
		{
			Name: "proper lookup_events",
			Spec: func() string {
				var input CustomCloudtrailLookupEventsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"lookup_events":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "lookup_events.NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomCloudtrailLookupEventsInput
				require.NoError(t, faker.FakeObject(&input))
				return `{"lookup_events":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
	})
}
