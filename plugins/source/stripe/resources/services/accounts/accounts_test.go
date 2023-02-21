package accounts_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/accounts"
)

func TestAccounts(t *testing.T) {
	client.MockTestHelper(t, accounts.Accounts(), client.TestOptions{})
}
