package tableoptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cloudquery/plugin-sdk/faker"
)

func TestAAListFindings(t *testing.T) {
	u := CustomAccessAnalyzerListFindingsInput{}
	require.NoError(t, faker.FakeObject(&u))

	api := AccessanalyzerFindings{
		ListFindingOpts: []CustomAccessAnalyzerListFindingsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")
	api.ListFindingOpts[0].NextToken = nil

	err = api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set AnalyzerARN in ListFindings")
	api.ListFindingOpts[0].AnalyzerArn = nil

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	err = api.Validate()
	assert.Nil(t, err)
}
