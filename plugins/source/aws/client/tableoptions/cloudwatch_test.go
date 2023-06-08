package tableoptions

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/stretchr/testify/assert"
)

func TestCloudwatchListMetrics(t *testing.T) {
	u := CustomCloudwatchListMetricsInput{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	api := CloudwatchAPIs{
		ListMetricsOpts: []CustomCloudwatchListMetricsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListMetrics")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListMetricsOpts[0].NextToken = nil
	err = api.Validate()
	assert.Nil(t, err)
}
