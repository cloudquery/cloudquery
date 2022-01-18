//go:build integration
// +build integration

package kms

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationKmsKeys(t *testing.T) {
	client.AWSTestHelper(t, KmsKeys())
}
