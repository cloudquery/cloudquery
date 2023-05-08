package table_options

import (
	"reflect"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_option_inputs/accessanalyzer_input"
	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/plugin-sdk/faker"
)

func TestAAListFindings(t *testing.T) {
	u := accessanalyzer_input.ListFindingsInput{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}

	api := AaFindings{
		ListFindingOpts: u,
	}
	// Ensure that the validation works as expected
	_, err := api.ListFindings()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")
	api.ListFindingOpts.NextToken = nil

	_, err = api.ListFindings()
	assert.EqualError(t, err, "invalid input: cannot set AnalyzerARN in ListFindings")
	api.ListFindingOpts.AnalyzerArn = nil

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	input, err := api.ListFindings()
	nilFields := findNilOrDefaultFields(reflect.ValueOf(*input), []string{})

	assert.Equal(t, nilFields, []string{"AnalyzerArn", "NextToken"}, "These are the only fields that should have a default value")
	assert.Nil(t, err)
}
