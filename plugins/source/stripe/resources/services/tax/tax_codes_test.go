package tax_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/tax"
)

func TestTaxCodes(t *testing.T) {
	client.MockTestHelper(t, tax.TaxCodes(), client.TestOptions{})
}
