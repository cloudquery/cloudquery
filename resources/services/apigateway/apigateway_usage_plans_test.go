// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayUsagePlans(t *testing.T) {
	client.AWSTestHelper(t, ApigatewayUsagePlans(),
		"./snapshots")
}
