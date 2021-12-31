// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEc2Images(t *testing.T) {
	client.AWSTestHelper(t, Ec2Images())
}
