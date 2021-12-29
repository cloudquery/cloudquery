// +build integration

package wafv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationWAFv2ManagedRuleGroups(t *testing.T) {
	client.AWSTestHelper(t, Wafv2ManagedRuleGroups(),
		"./snapshots")
}
