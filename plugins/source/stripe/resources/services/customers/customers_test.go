package customers_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/customers"
)

func TestCustomers(t *testing.T) {
	client.MockTestHelper(t, customers.Customers(), client.TestOptions{})
}
