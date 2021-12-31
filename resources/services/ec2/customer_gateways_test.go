// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEc2CustomerGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2CustomerGateways())
}
