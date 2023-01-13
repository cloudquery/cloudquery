package payouts_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/payouts"
)

func TestPayouts(t *testing.T) {
	client.MockTestHelper(t, payouts.Payouts(), client.TestOptions{})
}
