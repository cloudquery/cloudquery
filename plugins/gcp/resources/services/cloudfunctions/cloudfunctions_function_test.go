//go:build integration
// +build integration

package cloudfunctions

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationCloudFunctionFunction(t *testing.T) {
	client.GcpTestHelper(t, CloudfunctionsFunction())
}
