//go:build integration
// +build integration

package eks

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEksClusters(t *testing.T) {
	client.AWSTestHelper(t, EksClusters())
}
