//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-digitalocean/client"
)

func TestIntegrationDomains(t *testing.T) {
	client.DOTestHelper(t, Domains())
}
