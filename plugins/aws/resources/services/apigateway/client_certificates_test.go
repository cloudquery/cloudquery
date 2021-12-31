// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayClientCertificates(t *testing.T) {
	client.AWSTestHelper(t, ApigatewayClientCertificates())
}
