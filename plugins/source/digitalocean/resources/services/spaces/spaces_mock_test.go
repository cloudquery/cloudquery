//go:build integration
// +build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

//todo fix it

func TestIntegrationSpaces(t *testing.T) {
	client.DOTestHelper(t, Spaces())
}
