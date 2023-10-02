package tableoptions

import (
	"encoding/json"
	"strings"
	"testing"

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

func TestLookupEventsJSONSchema(t *testing.T) {
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_cloudtrail_events":{}}`,
		},
		{
			name: "proper",
			spec: func() string {
				var apis CloudtrailAPIs
				require.NoError(t, faker.FakeObject(&apis))
				require.Len(t, apis.LookupEventsOpts, 1)

				// remove prohibited fields
				apis.LookupEventsOpts[0].NextToken = nil

				data, err := json.MarshalIndent(TableOptions{CloudTrailEvents: &apis}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "NextToken is present",
			err:  true,
			spec: func() string {
				var apis CloudtrailAPIs
				require.NoError(t, faker.FakeObject(&apis))
				require.Len(t, apis.LookupEventsOpts, 1)

				data, err := json.MarshalIndent(TableOptions{CloudTrailEvents: &apis}, "", "  ")
				require.NoError(t, err)

				return string(data)
			}(),
		},
	})
}
