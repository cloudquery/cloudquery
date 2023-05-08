package table_options

import (
	"reflect"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_option_inputs/inspector2_input"
	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/plugin-sdk/faker"
)

func TestInspector2ListFindings(t *testing.T) {
	u := inspector2_input.ListFindingsInput{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}

	api := Inspector2APIs{
		ListFindingOpts: u,
	}
	// Ensure that the validation works as expected
	_, err := api.ListFindings()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListFindingOpts.NextToken = nil
	input, err := api.ListFindings()
	nilFields := findNilOrDefaultFields(reflect.ValueOf(*input), []string{})

	assert.Equal(t, nilFields, []string{"NextToken"}, "These are the only fields that should have a default value")
	assert.Nil(t, err)
}
