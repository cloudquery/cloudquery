// +build integration

package dms

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationDmsReplicationInstances(t *testing.T) {
	client.AWSTestHelper(t, DmsReplicationInstances(),
		"./snapshots")
}
