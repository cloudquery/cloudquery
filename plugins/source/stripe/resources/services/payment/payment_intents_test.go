package payment_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/payment"
)

func TestPaymentIntents(t *testing.T) {
	client.MockTestHelper(t, payment.PaymentIntents(), client.TestOptions{})
}
