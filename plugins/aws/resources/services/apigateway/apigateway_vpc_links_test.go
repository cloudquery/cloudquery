// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayVpcLinks(t *testing.T) {
	client.AWSTestHelper(t, ApigatewayVpcLinks(),
		"./snapshots")
}
