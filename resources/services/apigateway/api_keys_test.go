//go:build integration
// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayAPIKeys(t *testing.T) {
	client.AWSTestHelper(t, ApigatewayAPIKeys())
}
