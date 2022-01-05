//go:build integration
// +build integration

package resource_manager

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationResourceManagerProjects(t *testing.T) {
	client.GcpTestHelper(t, ResourceManagerProjects())
}
