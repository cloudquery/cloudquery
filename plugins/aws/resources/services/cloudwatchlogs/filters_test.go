//go:build integration
// +build integration

package cloudwatchlogs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCloudwatchlogsFilters(t *testing.T) {
	client.AWSTestHelper(t, CloudwatchlogsFilters())
}
