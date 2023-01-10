package balance_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/balance"
)

func TestBalance(t *testing.T) {
	client.MockTestHelper(t, balance.Balance(), client.TestOptions{})
}
