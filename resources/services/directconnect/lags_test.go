// +build integration

package directconnect

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationDirectconnectLags(t *testing.T) {
	client.AWSTestHelper(t, DirectconnectLags())
}
