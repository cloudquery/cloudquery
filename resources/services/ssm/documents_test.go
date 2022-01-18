//go:build integration
// +build integration

package ssm

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSSMDocuments(t *testing.T) {
	client.AWSTestHelper(t, SsmDocuments())
}
