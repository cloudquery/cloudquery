// +build integration

package autoscaling

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationAutoscalingGroups(t *testing.T) {
	client.AWSTestHelper(t, AutoscalingGroups(),
		"./snapshots")
}
