// +build integration

package route53

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRoute53HostedZones(t *testing.T) {
	client.AWSTestHelper(t, Route53HostedZones())
}
