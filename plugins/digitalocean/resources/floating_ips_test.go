//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/digitalocean/client"
)

func TestIntegrationFloatingIps(t *testing.T) {
	client.DOTestHelper(t, FloatingIps())
}
