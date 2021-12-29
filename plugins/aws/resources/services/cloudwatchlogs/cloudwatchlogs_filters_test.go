// +build integration

package cloudwatchlogs

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationCloudwatchlogsFilters(t *testing.T) {
	client.AWSTestHelper(t, CloudwatchlogsFilters(),
		"./snapshots")
}
