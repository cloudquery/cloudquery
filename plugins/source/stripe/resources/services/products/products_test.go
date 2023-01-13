package products_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/products"
)

func TestProducts(t *testing.T) {
	client.MockTestHelper(t, products.Products(), client.TestOptions{})
}
