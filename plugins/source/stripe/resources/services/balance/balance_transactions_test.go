package balance_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/balance"
)

func TestBalanceTransactions(t *testing.T) {
	client.MockTestHelper(t, balance.BalanceTransactions(), client.TestOptions{})
}
