package tableoptions

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/stretchr/testify/assert"
)

func TestLookupEvents(t *testing.T) {
	u := CustomLookupEventsOpts{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
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
