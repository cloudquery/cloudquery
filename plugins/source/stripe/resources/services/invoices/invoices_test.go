package invoices_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/invoices"
)

func TestInvoices(t *testing.T) {
	client.MockTestHelper(t, invoices.Invoices(), client.TestOptions{})
}
