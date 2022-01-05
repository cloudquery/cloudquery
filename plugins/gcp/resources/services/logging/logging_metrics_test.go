//go:build integration
// +build integration

package logging

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationLoggingMetrics(t *testing.T) {
	client.GcpTestHelper(t, LoggingMetrics())
}
