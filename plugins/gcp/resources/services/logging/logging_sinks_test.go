//go:build integration
// +build integration

package logging

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationLoggingSinks(t *testing.T) {
	client.GcpTestHelper(t, LoggingSinks())
}
