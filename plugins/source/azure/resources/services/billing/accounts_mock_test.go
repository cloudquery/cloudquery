package billing

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/gorilla/mux"
)

func createAccounts(router *mux.Router) error {
	// BillingAccounts are mocked in `MockTestHelper` since those are fetched during initialization
	return nil
}

func TestAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccounts)
}
