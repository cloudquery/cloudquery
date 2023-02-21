package subscription_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/subscription"
)

func TestSubscriptionSchedules(t *testing.T) {
	client.MockTestHelper(t, subscription.SubscriptionSchedules(), client.TestOptions{})
}
