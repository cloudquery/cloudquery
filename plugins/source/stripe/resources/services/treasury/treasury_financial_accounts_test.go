package treasury_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/treasury"
)

func TestTreasuryFinancialAccounts(t *testing.T) {
	client.MockTestHelper(t, treasury.TreasuryFinancialAccounts(), client.TestOptions{})
}
