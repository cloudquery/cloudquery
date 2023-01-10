package invoices_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/invoices"
)

func TestInvoiceItems(t *testing.T) {
	client.MockTestHelper(t, invoices.InvoiceItems(), client.TestOptions{})
}
