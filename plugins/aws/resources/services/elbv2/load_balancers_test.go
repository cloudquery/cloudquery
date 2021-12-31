// +build integration

package elbv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationElbv2LoadBalancers(t *testing.T) {
	client.AWSTestHelper(t, Elbv2LoadBalancers())
}
