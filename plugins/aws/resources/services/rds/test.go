//go:build integration
// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRdsInstances(t *testing.T) {
	client.AWSTestHelper(t, RdsInstances())
}

func TestIntegrationRdsSubnetGroups(t *testing.T) {
	client.AWSTestHelper(t, RdsSubnetGroups())
}

func TestIntegrationRdsClusters(t *testing.T) {
	client.AWSTestHelper(t, RdsClusters())
}
