//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-digitalocean/client"
)

func TestIntegrationDatabases(t *testing.T) {
	client.DOTestHelper(t, Databases())
}
