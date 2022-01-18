//go:build integration
// +build integration

package apigatewayv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayv2VpcLinks(t *testing.T) {
	client.AWSTestHelper(t, Apigatewayv2VpcLinks())
}
