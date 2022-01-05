//go:build integration
// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationIamRoles(t *testing.T) {
	client.GcpTestHelper(t, IamRoles())
}
