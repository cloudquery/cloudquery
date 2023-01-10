package charges_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/charges"
)

func TestCharges(t *testing.T) {
	client.MockTestHelper(t, charges.Charges(), client.TestOptions{})
}
