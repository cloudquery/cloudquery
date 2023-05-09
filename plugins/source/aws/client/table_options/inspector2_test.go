package table_options

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/plugin-sdk/faker"
)

func TestInspector2ListFindings(t *testing.T) {
	u := CustomInspector2ListFindingsInput{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}

	api := Inspector2APIs{
		ListFindingOpts: u,
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListFindingOpts.NextToken = nil
	err = api.Validate()
	
	assert.Nil(t, err)
}
