//go:build integration
// +build integration

package wafv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationWAFv2WebACLs(t *testing.T) {
	client.AWSTestHelper(t, Wafv2WebAcls())
}
