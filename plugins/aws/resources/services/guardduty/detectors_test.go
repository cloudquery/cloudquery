// +build integration

package guardduty

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationGuarddutyDetectors(t *testing.T) {
	client.AWSTestHelper(t, GuarddutyDetectors())
}
