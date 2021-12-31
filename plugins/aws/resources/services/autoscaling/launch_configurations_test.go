// +build integration

package autoscaling

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationAutoscalingLaunchConfigurations(t *testing.T) {
	client.AWSTestHelper(t,
		AutoscalingLaunchConfigurations())
}
