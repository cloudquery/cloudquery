// +build integration

package apigatewayv2

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationApigatewayv2ApisTest(t *testing.T) {
	client.AWSTestHelper(t, Apigatewayv2Apis(),
		"./snapshots")
}
