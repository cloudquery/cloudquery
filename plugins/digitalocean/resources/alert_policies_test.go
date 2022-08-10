//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/digitalocean/client"
)

func TestIntegrationAlertPolicies(t *testing.T) {
	client.DOTestHelper(t, AlertPolicies())
}
