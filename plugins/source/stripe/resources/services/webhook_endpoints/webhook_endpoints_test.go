package webhook_endpoints_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/webhook_endpoints"
)

func TestWebhookEndpoints(t *testing.T) {
	client.MockTestHelper(t, webhook_endpoints.WebhookEndpoints(), client.TestOptions{})
}
