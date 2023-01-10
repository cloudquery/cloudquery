package checkout_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/checkout"
)

func TestCheckoutSessions(t *testing.T) {
	client.MockTestHelper(t, checkout.CheckoutSessions(), client.TestOptions{})
}
