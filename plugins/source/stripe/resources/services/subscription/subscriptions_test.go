package subscription_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/subscription"
)

func TestSubscriptions(t *testing.T) {
	client.MockTestHelper(t, subscription.Subscriptions(), client.TestOptions{})
}
