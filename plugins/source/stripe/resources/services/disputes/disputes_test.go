package disputes_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/disputes"
)

func TestDisputes(t *testing.T) {
	client.MockTestHelper(t, disputes.Disputes(), client.TestOptions{})
}
