//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

func TestIntegrationBillingHistory(t *testing.T) {
	client.DOTestHelper(t, BillingHistory())
}
