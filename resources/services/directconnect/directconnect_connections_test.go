// +build integration

package directconnect

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationDirectconnectConnections(t *testing.T) {
	client.AWSTestHelper(t, DirectconnectConnections(),
		"./snapshots")
}
