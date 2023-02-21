package refunds_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/refunds"
)

func TestRefunds(t *testing.T) {
	client.MockTestHelper(t, refunds.Refunds(), client.TestOptions{})
}
