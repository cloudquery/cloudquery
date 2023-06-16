package billing

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/gorilla/mux"
)

func createPeriods(router *mux.Router) error {
	// BillingPeriods are mocked in `MockTestHelper` since those are fetched during initialization
	return nil
}

func TestPeriods(t *testing.T) {
	client.MockTestHelper(t, Periods(), createPeriods)
}
