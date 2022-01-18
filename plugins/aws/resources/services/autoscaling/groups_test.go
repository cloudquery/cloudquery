//go:build integration
// +build integration

package autoscaling

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationAutoscalingGroups(t *testing.T) {
	client.AWSTestHelper(t, AutoscalingGroups())
}
