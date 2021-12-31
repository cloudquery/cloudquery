// +build integration

package directconnect

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationDirectConnectGateways(t *testing.T) {
	client.AWSTestHelper(t, DirectconnectGateways())
}
