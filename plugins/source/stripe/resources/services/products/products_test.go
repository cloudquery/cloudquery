package products_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/products"
)

func TestProductsProduct(t *testing.T) {
	tbl := products.Products()
	if err := tbl.Transform(tbl); err != nil {
		t.Fatal(err)
	}
	for i, c := range tbl.Columns {
		if c.Name == "attributes" || c.Name == "deactivate_on" {
			tbl.Columns[i].IgnoreInTests = true
		}
	}
	client.MockTestHelper(t, tbl)
}
