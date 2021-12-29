// +build integration

package ecs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEcsClusters(t *testing.T) {
	client.AWSTestHelper(t, EcsClusters(),
		"./snapshots")
}
