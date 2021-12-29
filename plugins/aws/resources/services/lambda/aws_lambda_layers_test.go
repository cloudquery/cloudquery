// +build integration

package lambda

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationLambdaLayers(t *testing.T) {
	client.AWSTestHelper(t, LambdaLayers(),
		"./snapshots")
}
