// +build integration

package apigatewayv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayv2ApisTest(t *testing.T) {
	client.AWSTestHelper(t, Apigatewayv2Apis())
}
