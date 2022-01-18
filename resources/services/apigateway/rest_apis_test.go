//go:build integration
// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	client.AWSTestHelper(t, ApigatewayRestApis())
}
