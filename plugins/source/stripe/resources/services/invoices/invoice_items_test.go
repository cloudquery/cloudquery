package invoices_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/invoices"
)

func TestInvoicesInvoiceItem(t *testing.T) {
	tbl := invoices.InvoiceItems()
	if err := tbl.Transform(tbl); err != nil {
		t.Fatal(err)
	}
	for i, c := range tbl.Columns {
		if c.Name == "plan" {
			tbl.Columns[i].IgnoreInTests = true
		}
	}
	client.MockTestHelper(t, tbl)
}
