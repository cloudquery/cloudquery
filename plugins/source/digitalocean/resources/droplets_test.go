//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

func TestIntegrationDroplets(t *testing.T) {
	client.DOTestHelper(t, Droplets())
}
