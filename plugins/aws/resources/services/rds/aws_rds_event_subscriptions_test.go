// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRdsEventSubscriptions(t *testing.T) {
	client.AWSTestHelper(t, RdsEventSubscriptions(),
		"./snapshots")
}
