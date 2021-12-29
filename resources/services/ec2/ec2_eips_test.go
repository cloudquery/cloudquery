// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEc2Eips(t *testing.T) {
	client.AWSTestHelper(t, Ec2Eips(),
		"./snapshots")
}
