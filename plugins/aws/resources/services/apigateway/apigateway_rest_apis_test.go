// +build integration

package apigateway

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	client.AWSTestHelper(t, ApigatewayRestApis(), "./snapshots")
}
