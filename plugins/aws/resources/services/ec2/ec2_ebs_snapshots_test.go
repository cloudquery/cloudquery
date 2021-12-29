// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEc2EbsSnapshots(t *testing.T) {
	client.AWSTestHelper(t, Ec2EbsSnapshots(),
		"./snapshots")
}
