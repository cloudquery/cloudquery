package topups_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/topups"
)

func TestTopups(t *testing.T) {
	client.MockTestHelper(t, topups.Topups(), client.TestOptions{})
}
