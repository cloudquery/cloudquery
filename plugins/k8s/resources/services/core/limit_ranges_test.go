//go:build integration
// +build integration

package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/k8s/client"
)

func TestIntegrationLimitRanges(t *testing.T) {
	client.K8sTestHelper(t, LimitRanges(), "./snapshots")
}
