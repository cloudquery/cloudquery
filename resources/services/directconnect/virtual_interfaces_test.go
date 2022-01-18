//go:build integration
// +build integration

package directconnect

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationDirectconnectVirtualInterfaces(t *testing.T) {
	client.AWSTestHelper(t, DirectconnectVirtualInterfaces())
}
