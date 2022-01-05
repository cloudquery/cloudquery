//go:build integration
// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationIamServiceAccounts(t *testing.T) {
	client.GcpTestHelper(t, IamServiceAccounts())
}
