//go:build integration
// +build integration

package cloudwatch

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCloudwatchAlarms(t *testing.T) {
	client.AWSTestHelper(t, CloudwatchAlarms())
}
