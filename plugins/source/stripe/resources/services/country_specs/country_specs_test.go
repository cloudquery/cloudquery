package country_specs_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/country_specs"
)

func TestCountrySpecs(t *testing.T) {
	client.MockTestHelper(t, country_specs.CountrySpecs(), client.TestOptions{})
}
