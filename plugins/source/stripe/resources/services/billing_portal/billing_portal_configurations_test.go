package billing_portal_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/billing_portal"
)

func TestBillingPortalConfigurations(t *testing.T) {
	client.MockTestHelper(t, billing_portal.BillingPortalConfigurations(), client.TestOptions{})
}
