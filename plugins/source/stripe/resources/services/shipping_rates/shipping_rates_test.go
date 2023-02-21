package shipping_rates_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/shipping_rates"
)

func TestShippingRates(t *testing.T) {
	client.MockTestHelper(t, shipping_rates.ShippingRates(), client.TestOptions{})
}
