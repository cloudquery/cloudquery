package issuing_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/issuing"
)

func TestIssuingTransactions(t *testing.T) {
	client.MockTestHelper(t, issuing.IssuingTransactions(), client.TestOptions{})
}
