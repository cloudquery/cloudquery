//go:build integration
// +build integration

package waf

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationWAFWebACLs(t *testing.T) {
	client.AWSTestHelper(t, WafWebAcls())
}
