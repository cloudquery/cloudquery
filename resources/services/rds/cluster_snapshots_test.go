//go:build integration
// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRdsClusterSnapshots(t *testing.T) {
	client.AWSTestHelper(t, RdsClusterSnapshots())
}
