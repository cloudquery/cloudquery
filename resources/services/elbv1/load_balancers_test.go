//go:build integration
// +build integration

package elbv1

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationElbv1LoadBalancers(t *testing.T) {
	client.AWSTestHelper(t, Elbv1LoadBalancers())
}
