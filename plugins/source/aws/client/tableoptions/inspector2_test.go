package tableoptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cloudquery/plugin-sdk/faker"
)

func TestInspector2ListFindings(t *testing.T) {
	u := CustomInspector2ListFindingsInput{}
	require.NoError(t, faker.FakeObject(&u))

	api := Inspector2APIs{
		ListFindingsOpts: []CustomInspector2ListFindingsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListFindingsOpts[0].NextToken = nil
	err = api.Validate()

	assert.Nil(t, err)
}
