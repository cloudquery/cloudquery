//go:build integration
// +build integration

package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
)

func TestIntegrationNamespaces(t *testing.T) {
	client.K8sTestHelper(t, Namespaces(), "./snapshots")
}
