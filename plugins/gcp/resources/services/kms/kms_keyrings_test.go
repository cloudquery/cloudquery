//go:build integration
// +build integration

package kms

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationKmsKeyrings(t *testing.T) {
	client.GcpTestHelper(t, KmsKeyrings())
}
