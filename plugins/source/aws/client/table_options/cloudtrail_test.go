package table_options

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_options/inputs/cloudtrail_input"
	"github.com/cloudquery/plugin-sdk/faker"
)

func TestLookupEvents(t *testing.T) {
	u := cloudtrail_input.LookupEventsInput{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	api := CtAPIs{
		LookupEventsOpts: u,
	}
	// Ensure that the validation works as expected
	_, err := api.LookupEvents()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in LookupEvents")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.LookupEventsOpts.NextToken = nil
	input, err := api.LookupEvents()
	nilFields := findNilOrDefaultFields(reflect.ValueOf(*input), []string{})

	assert.Equal(t, nilFields, []string{"NextToken"}, "These are the only fields that should have a default value")
	assert.Nil(t, err)
}
