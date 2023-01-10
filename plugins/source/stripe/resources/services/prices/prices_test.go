package prices_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/prices"
)

func TestPrices(t *testing.T) {
	client.MockTestHelper(t, prices.Prices(), client.TestOptions{})
}
