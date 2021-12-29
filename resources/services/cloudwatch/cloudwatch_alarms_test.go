// +build integration

package cloudwatch

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationCloudwatchAlarms(t *testing.T) {
	client.AWSTestHelper(t, CloudwatchAlarms(),
		"./snapshots")
}
