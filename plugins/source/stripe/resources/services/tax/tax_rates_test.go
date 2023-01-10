package tax_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/tax"
)

func TestTaxRates(t *testing.T) {
	client.MockTestHelper(t, tax.TaxRates(), client.TestOptions{})
}
