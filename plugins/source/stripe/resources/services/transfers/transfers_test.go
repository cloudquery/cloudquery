package transfers_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/transfers"
)

func TestTransfers(t *testing.T) {
	client.MockTestHelper(t, transfers.Transfers(), client.TestOptions{})
}
