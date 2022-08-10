//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/digitalocean/client"
)

func TestIntegrationAccount(t *testing.T) {
	client.DOTestHelper(t, Account())
}
