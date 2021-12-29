// +build integration

package lambda

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationLambdaFunctions(t *testing.T) {
	client.AWSTestHelper(t, LambdaFunctions(),
		"./snapshots")
}
